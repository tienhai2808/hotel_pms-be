package usecase

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/InstayPMS/backend/internal/application/dto"
	"github.com/InstayPMS/backend/internal/application/port"
	"github.com/InstayPMS/backend/internal/domain/model"
	"github.com/InstayPMS/backend/internal/domain/repository"
	"github.com/InstayPMS/backend/internal/infrastructure/config"
	"github.com/InstayPMS/backend/pkg/errors"
	"github.com/InstayPMS/backend/pkg/utils"
	"github.com/sony/sonyflake/v2"
	"go.uber.org/zap"
)

type authUseCaseImpl struct {
	cfg       config.JWTConfig
	log       *zap.Logger
	idGen     *sonyflake.Sonyflake
	jwtPro    port.JWTProvider
	cachePro  port.CacheProvider
	userRepo  repository.UserRepository
	tokenRepo repository.TokenRepository
}

func NewAuthUseCase(
	cfg config.JWTConfig,
	log *zap.Logger,
	idGen *sonyflake.Sonyflake,
	jwtPro port.JWTProvider,
	cachePro port.CacheProvider,
	userRepo repository.UserRepository,
	tokenRepo repository.TokenRepository,
) AuthUseCase {
	return &authUseCaseImpl{
		cfg,
		log,
		idGen,
		jwtPro,
		cachePro,
		userRepo,
		tokenRepo,
	}
}

func (u *authUseCaseImpl) Login(ctx context.Context, ua string, req dto.LoginRequest) (*model.User, string, string, error) {
	user, err := u.userRepo.FindByUsernameWithOutletAndDepartment(ctx, req.Username)
	if err != nil {
		u.log.Error("find user by username failed", zap.String("username", req.Username), zap.Error(err))
		return nil, "", "", err
	}

	if user == nil {
		return nil, "", "", errors.ErrLoginFailed
	}

	if !user.IsActive {
		return nil, "", "", errors.ErrLoginFailed
	}

	if err = utils.VerifyPassword(req.Password, user.Password); err != nil {
		return nil, "", "", errors.ErrLoginFailed
	}

	redisKey := fmt.Sprintf("user_version:%s", strconv.Itoa(int(user.ID)))
	tokenVersion, err := u.cachePro.GetInt(ctx, redisKey)
	if err != nil {
		u.log.Error("get token version failed", zap.Error(err))
		return nil, "", "", err
	}

	if tokenVersion == 0 {
		if err = u.cachePro.SetInt(ctx, redisKey, 1, 0); err != nil {
			u.log.Error("save token version failed", zap.Error(err))
			return nil, "", "", err
		}
		tokenVersion = 1
	}

	accessToken, err := u.jwtPro.GenerateToken(user.ID, tokenVersion, u.cfg.AccessExpiresIn)
	if err != nil {
		u.log.Error("generate access token failed", zap.Error(err))
		return nil, "", "", err
	}

	refreshToken, hashedToken, err := utils.GenerateRefreshToken()
	if err != nil {
		u.log.Error("generate refresh token failed", zap.Error(err))
		return nil, "", "", err
	}

	id, err := u.idGen.NextID()
	if err != nil {
		u.log.Error("generate token id failed", zap.Error(err))
		return nil, "", "", err
	}

	token := &model.Token{
		ID:        id,
		UserID:    user.ID,
		Token:     hashedToken,
		UserAgent: utils.ConvertUserAgent(ua),
		RevokedAt: nil,
		ExpiresAt: time.Now().Add(u.cfg.RefreshExpiresIn),
	}

	if err := u.tokenRepo.Create(ctx, token); err != nil {
		u.log.Error("save token to database failed", zap.Error(err))
		return nil, "", "", err
	}

	return user, accessToken, refreshToken, nil
}

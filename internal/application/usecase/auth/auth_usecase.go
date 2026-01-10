package usecase

import (
	"context"
	"time"

	"github.com/InstayPMS/backend/internal/application/dto"
	"github.com/InstayPMS/backend/internal/domain/model"
)

type AuthUseCase interface {
	Login(ctx context.Context, ua string, req dto.LoginRequest) (*model.User, string, string, error)

	Logout(ctx context.Context, accessToken, refreshToken string, accessTTL time.Duration) error

	RefreshToken(ctx context.Context, ua, refreshToken string) (string, string, error)

	GetMe(ctx context.Context, userID int64) (*model.User, error)

	ChangePassword(ctx context.Context, userID int64, req dto.ChangePasswordRequest) error

	ForgotPassword(ctx context.Context, email string) (string, error)

	VerifyForgotPassword(ctx context.Context, req dto.VerifyForgotPasswordRequest) (string, error)

	ResetPassword(ctx context.Context, req dto.ResetPasswordRequest) error

	UpdateInfo(ctx context.Context, userID int64, req dto.UpdateInfoRequest) (*model.User, error)
}

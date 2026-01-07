package usecase

import (
	"github.com/sony/sonyflake/v2"
	"go.uber.org/zap"
)

type authUseCaseImpl struct {
	log   *zap.Logger
	idGen *sonyflake.Sonyflake
}

func NewAuthUseCase(
	log *zap.Logger,
	idGen *sonyflake.Sonyflake,
) AuthUseCase {
	return &authUseCaseImpl{
		log,
		idGen,
	}
}

func (u *authUseCaseImpl) Login() {}

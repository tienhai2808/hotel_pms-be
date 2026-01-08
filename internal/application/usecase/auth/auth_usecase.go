package usecase

import (
	"context"

	"github.com/InstayPMS/backend/internal/application/dto"
	"github.com/InstayPMS/backend/internal/domain/model"
)

type AuthUseCase interface {
	Login(ctx context.Context, ua string, req dto.LoginRequest) (*model.User, string, string, error)
}
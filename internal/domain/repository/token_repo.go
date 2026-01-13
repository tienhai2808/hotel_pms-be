package repository

import (
	"context"

	"github.com/InstaySystem/is_v2-be/internal/domain/model"
	"gorm.io/gorm"
)

type TokenRepository interface {
	Create(ctx context.Context, token *model.Token) error

	UpdateByToken(ctx context.Context, token string, updateData map[string]any) error

	FindByToken(ctx context.Context, token string) (*model.Token, error)

	UpdateAllByUserIDTx(tx *gorm.DB, userID int64, updateData map[string]any) error

	DeleteAllByUserIDTx(tx *gorm.DB, userID int64) error

	DeleteAllByUserIDsTx(tx *gorm.DB, userIDs []int64) error

	DeleteAllExpired(ctx context.Context) (int64, error)
}

package handler

import (
	authUC "github.com/InstayPMS/backend/internal/application/usecase/auth"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authUC authUC.AuthUseCase
}

func NewAuthHandler(authUC authUC.AuthUseCase) *AuthHandler {
	return &AuthHandler{authUC}
}

func (h *AuthHandler) Login(c *gin.Context) {
	
}
package handler

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/InstayPMS/backend/internal/application/dto"
	authUC "github.com/InstayPMS/backend/internal/application/usecase/auth"
	"github.com/InstayPMS/backend/internal/infrastructure/config"
	"github.com/InstayPMS/backend/pkg/constants"
	"github.com/InstayPMS/backend/pkg/errors"
	"github.com/InstayPMS/backend/pkg/mapper"
	"github.com/InstayPMS/backend/pkg/utils"
	"github.com/InstayPMS/backend/pkg/validator"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	cfg    *config.Config
	authUC authUC.AuthUseCase
}

func NewAuthHandler(
	cfg *config.Config,
	authUC authUC.AuthUseCase,
) *AuthHandler {
	return &AuthHandler{
		cfg,
		authUC,
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		field, tag, param := validator.HandleRequestError(err)
		c.Error(errors.ErrBadRequest.WithData(gin.H{
			"field": field,
			"tag":   tag,
			"param": param,
		}))
		return
	}

	user, accessToken, refreshToken, err := h.authUC.Login(ctx, c.Request.UserAgent(), req)
	if err != nil {
		c.Error(err)
		return
	}

	isSecure := c.Request.TLS != nil || c.GetHeader("X-Forwarded-Proto") == "https"
	domain := utils.ExtractRootDomain(c.Request.Host)

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie(h.cfg.JWT.AccessName, accessToken, int(h.cfg.JWT.AccessExpiresIn.Seconds()), "/", domain, isSecure, true)
	c.SetCookie(h.cfg.JWT.RefreshName, refreshToken, int(h.cfg.JWT.RefreshExpiresIn.Seconds()), fmt.Sprintf("%s/auth/refresh-token", h.cfg.Server.APIPrefix), domain, isSecure, true)

	utils.APIResponse(c, http.StatusOK, constants.CodeLoginSuccess, "Login successfully", gin.H{
		"user": mapper.ToUserResponse(user),
	})
}

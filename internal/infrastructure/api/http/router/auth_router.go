package router

import (
	"github.com/InstayPMS/backend/internal/infrastructure/api/http/handler"
	"github.com/InstayPMS/backend/internal/infrastructure/api/http/middleware"
	"github.com/gin-gonic/gin"
)

func (r *Router) setupAuthRoutes(rg *gin.RouterGroup, authMid *middleware.AuthMiddleware, hdl *handler.AuthHandler) {
	auth := rg.Group("/auth")
	{
		auth.POST("/login", hdl.Login)

		auth.POST("/logout", authMid.IsAuthentication(), authMid.AttachTokens(), hdl.Logout)

		auth.POST("/refresh-token", hdl.RefreshToken)

		auth.GET("/me", authMid.IsAuthentication(), hdl.GetMe)

		auth.POST("/change-password", authMid.IsAuthentication(), hdl.ChangePassword)

		auth.POST("/forgot-password", hdl.ForgotPassword)

		auth.POST("/forgot-password/verify", hdl.VerifyForgotPassword)

		auth.POST("/reset-password", hdl.ResetPassword)

		auth.POST("/update-info", authMid.IsAuthentication(), hdl.UpdateInfo)
	}
}

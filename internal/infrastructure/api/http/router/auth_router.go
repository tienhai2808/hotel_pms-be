package router

import (
	"github.com/InstayPMS/backend/internal/infrastructure/api/http/handler"
	"github.com/gin-gonic/gin"
)

func (r *Router) setupAuthRoutes(rg *gin.RouterGroup, hdl *handler.AuthHandler) {
	auth := rg.Group("/auth")
	{
		auth.POST("/login", hdl.Login)
	}
}

package mapper

import "github.com/gin-gonic/gin"

func ToAPIResponse(c *gin.Context, statusCode, internalCode int, message string, data any)

package handler

import (
	fileUC "github.com/InstayPMS/backend/internal/application/usecase/file"
	"github.com/gin-gonic/gin"
)

type FileHandler struct {
	fileUC fileUC.FileUseCase
}

func NewFileHandler(fileUC fileUC.FileUseCase) *FileHandler {
	return &FileHandler{fileUC}
}

func (h *FileHandler) UploadPresignedURLs(c *gin.Context) {
	
}

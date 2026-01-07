package di

import "github.com/InstayPMS/backend/internal/infrastructure/api/http/handler"

func (c *Container) initHandlers() {
	c.FileHandler = handler.NewFileHandler(c.FileUseCase)
}
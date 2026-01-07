package container

import "github.com/InstayPMS/backend/internal/infrastructure/api/http/handler"

func (c *Container) initHandlers() {
	c.FileHdl = handler.NewFileHandler(c.FileUC)
	c.AuthHdl = handler.NewAuthHandler(c.AuthUC)
}

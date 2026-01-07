package container

import (
	authUC "github.com/InstayPMS/backend/internal/application/usecase/auth"
	fileUC "github.com/InstayPMS/backend/internal/application/usecase/file"
)

func (c *Container) initUseCases() {
	c.FileUC = fileUC.NewFileUseCase(c.Cfg, c.Stor, c.Log)
	c.AuthUC = authUC.NewAuthUseCase(c.Log, c.IDGen)
}

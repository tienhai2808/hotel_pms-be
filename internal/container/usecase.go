package container

import (
	authUC "github.com/InstayPMS/backend/internal/application/usecase/auth"
	fileUC "github.com/InstayPMS/backend/internal/application/usecase/file"
)

func (c *Container) initUseCases() {
	c.fileUC = fileUC.NewFileUseCase(c.cfg, c.stor, c.Log)
	c.authUC = authUC.NewAuthUseCase(c.cfg.JWT, c.db.Gorm, c.Log, c.idGen, c.jwtPro, c.cachePro, c.MQPro, c.userRepo, c.tokenRepo)
}

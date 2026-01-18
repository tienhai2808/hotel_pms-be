package container

import (
	authUC "github.com/InstaySystem/is_v2-be/internal/application/usecase/auth"
	departmentUC "github.com/InstaySystem/is_v2-be/internal/application/usecase/department"
	fileUC "github.com/InstaySystem/is_v2-be/internal/application/usecase/file"
	userUC "github.com/InstaySystem/is_v2-be/internal/application/usecase/user"
	"github.com/InstaySystem/is_v2-be/internal/infrastructure/persistence/orm"
)

func (c *Container) initLogic() {
	c.UserRepo = orm.NewUserRepository(c.DB.ORM())
	c.TokenRepo = orm.NewTokenRepository(c.DB.ORM())
	c.departmentRepo = orm.NewDepartmentRepository(c.DB.ORM())

	c.fileUC = fileUC.NewFileUseCase(c.cfg.MinIO, c.stor.Client(), c.stor.Presigner(), c.Log.Logger())
	c.authUC = authUC.NewAuthUseCase(c.cfg.JWT, c.DB.ORM(), c.Log.Logger(), c.IDGen.Generator(), c.jwtPro, c.cachePro, c.MQPro, c.UserRepo, c.TokenRepo)
	c.userUC = userUC.NewUserUseCase(c.DB.ORM(), c.Log.Logger(), c.IDGen.Generator(), c.cachePro, c.UserRepo, c.departmentRepo, c.TokenRepo)
	c.departmentUC = departmentUC.NewDepartmentUseCase(c.Log.Logger(), c.IDGen.Generator(), c.departmentRepo)
}

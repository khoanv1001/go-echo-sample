package core

import (
	"github.com/khoanv1001/go-echo-sample/config"
	"github.com/khoanv1001/go-echo-sample/modules/core/handlers"
	"github.com/khoanv1001/go-echo-sample/modules/core/repositories"
	"github.com/khoanv1001/go-echo-sample/modules/core/usecases"
	"github.com/khoanv1001/go-echo-sample/pkg/middlewares"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

var Module ModuleInstance = &coreModule{}

type coreModule struct{}

func (coreModule) RegisterRepositories(container *dig.Container) error {
	container.Provide(repositories.NewPgsqlUserRepository)
	container.Provide(repositories.NewPgsqlOrgRepository)
	container.Provide(repositories.NewPgsqlUserOrgRepository)
	return nil
}

func (coreModule) RegisterUseCases(container *dig.Container) error {
	container.Provide(usecases.NewUserUsecase)
	container.Provide(usecases.NewOrgUsecase)
	return nil
}

func (coreModule) RegisterHandlers(g *echo.Group, container *dig.Container) error {
	return container.Invoke(func(
		appConf *config.AppConfig,
		middManager *middlewares.MiddlewareManager,
		userUsecase usecases.UserUsecase,
		orgUsecase usecases.OrgUsecase,
	) {
		handlers.NewOrgHandler(g, middManager, orgUsecase)
		handlers.NewUserHandler(g, middManager, userUsecase)
		handlers.NewAuthHandler(g, middManager, userUsecase, appConf)
	})
}

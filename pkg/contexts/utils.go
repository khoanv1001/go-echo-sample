package contexts

import (
	coreDomains "github.com/khoanv1001/go-echo-sample/modules/core/domains"
	projectDomains "github.com/khoanv1001/go-echo-sample/modules/projects/domains"
	"github.com/khoanv1001/go-echo-sample/pkg/constants"
	"github.com/labstack/echo/v4"
)

func GetUserFromContext(c echo.Context) (usr *coreDomains.UserWithRoles, err error) {
	uCtx := c.Get(constants.ContextKeyUser)
	usr, ok := uCtx.(*coreDomains.UserWithRoles)
	if !ok || usr == nil {
		err = constants.ErrUnauthorized
	}
	return
}

func GetOrgFromContext(c echo.Context) *coreDomains.Org {
	org := c.Get(constants.ContextKeyOrg)
	if org != nil {
		return org.(*coreDomains.Org)
	}
	return nil
}

func GetProjectFromContext(c echo.Context) *projectDomains.Project {
	proj := c.Get(constants.ContextKeyProject)
	if proj != nil {
		return proj.(*projectDomains.Project)
	}
	return nil
}

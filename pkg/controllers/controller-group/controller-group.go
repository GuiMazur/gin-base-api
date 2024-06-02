package controllergroup

import (
	"gin-base-api/pkg/config"
	"gin-base-api/pkg/modules/app"
	"gin-base-api/pkg/modules/token"
	"gin-base-api/pkg/modules/user"

	"gorm.io/gorm"
)

type ControllerGroup struct {
	*app.AppController
	*user.UserController
}

var controllerGroupInstance *ControllerGroup

func New(db *gorm.DB) *ControllerGroup {
	if controllerGroupInstance == nil {
		controllerGroupInstance = &ControllerGroup{
			AppController:        app.NewController(app.NewService()),
			UserController:    user.NewController(user.NewService(db, token.NewService(config.New()))),
		}
	}
	return controllerGroupInstance
}

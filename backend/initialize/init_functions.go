package initialize

import (
	"sister-backend/app/controller"
)

var InitFunctions = []InitFunc{
	controller.InitializeUserController,
	controller.InitializeAdminController,
}

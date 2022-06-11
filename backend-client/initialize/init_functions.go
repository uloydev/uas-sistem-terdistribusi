package initialize

import (
	"sister-backend/app/controller"
)

var InitFunctions = []InitFunc{
	controller.InitializeLiceseController,
	controller.InitializeProductController,
}

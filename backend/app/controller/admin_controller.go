package controller

import (
	"sister-backend/app/middleware"
	"sister-backend/app/model"
	"sister-backend/app/repository"
	"sister-backend/app/service"
	"sister-backend/exception"
	"sister-backend/utils"

	"github.com/go-gomail/gomail"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AdminController struct {
	Service  service.AdminService
	MailConn *gomail.Dialer
}

func NewAdminController(AdminService *service.AdminService, MailConn *gomail.Dialer) AdminController {
	return AdminController{Service: *AdminService, MailConn: MailConn}
}

func InitializeAdminController(api *fiber.Group, DB *gorm.DB, MailConn *gomail.Dialer) {
	adminRepo := repository.NewAdminRepository(DB)
	adminService := service.NewAdminService(adminRepo.(*repository.AdminRepository))
	adminController := NewAdminController(adminService.(*service.AdminService), MailConn)
	adminController.Route(api)
}

func (controller *AdminController) Route(api *fiber.Group) {
	api.Post("/admin", controller.Create)
	api.Get("/admin", middleware.JWTProtected(), controller.List)
	api.Post("/admin/auth", controller.GetNewAccessToken)
}

// CreateAdmin is a function to insert Admin to database
// @Security ApiKeyAuth
// @Summary      Create Admin
// @Description  Create New Admin / Register Admin
// @Tags         admin
// @Accept       json
// @Produce      json
// @Param        admin  body      model.AdminRequest  true  "Register Admin"
// @Success      200   {object}  model.WebResponse{data=model.AdminResponse}
// @Failure      500   {object}  model.WebResponse{data=string}
// @Failure      400   {object}  model.WebResponse{data=string}
// @Router       /v1/admin [post]
func (controller *AdminController) Create(c *fiber.Ctx) error {
	var request model.AdminRequest

	err := c.BodyParser(&request)
	exception.PanicWhenError(err)

	request.Password = utils.GenerateHash(request.Password)

	response := controller.Service.Create(request)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

// GetAllAdmin is a function to get all Admin data from database
// @Security ApiKeyAuth
// @Summary      Get All Admin
// @Description  get all Admin data from database
// @Tags         admin
// @Accept       json
// @Produce      json
// @Success      200   {object}  model.WebResponse{data=[]model.AdminResponse}
// @Failure      500   {object}  model.WebResponse{data=string}
// @Failure      400   {object}  model.WebResponse{data=string}
// @Router       /v1/admin [get]
func (controller *AdminController) List(c *fiber.Ctx) error {
	responses := controller.Service.List()

	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   responses,
	})
}

// AuthAdmin is a function to authenticate admin
// @Summary      Auth Admin
// @Description  authenticate admin / login admin
// @Tags         admin
// @Accept       json
// @Produce      json
// @Param        auth  body      model.AuthRequest  true  "Auth admin"
// @Success      200   {object}  model.WebResponse{data=model.AuthResponse}
// @Failure      500   {object}  model.WebResponse{data=string}
// @Failure      400   {object}  model.WebResponse{data=string}
// @Router       /v1/admin/auth [post]
func (controller *AdminController) GetNewAccessToken(c *fiber.Ctx) error {
	// Generate a new Access token.
	var request model.AuthRequest
	err := c.BodyParser(&request)
	exception.PanicWhenError(err)

	auth := controller.Service.FindByAuth(request)

	utils.ValidatePassword(request.Password, auth.Password)

	jwtToken := utils.GenerateNewToken(&auth, true)

	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "success",
		Data:   jwtToken,
	})
}

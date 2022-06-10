package controller

import (
	"sister-backend/app/middleware"
	"sister-backend/app/model"
	"sister-backend/app/repository"
	"sister-backend/app/service"
	"sister-backend/emails"
	"sister-backend/exception"
	"sister-backend/utils"

	"github.com/go-gomail/gomail"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserController struct {
	Service  service.UserService
	MailConn *gomail.Dialer
}

func NewUserController(UserService *service.UserService, MailConn *gomail.Dialer) BaseController {
	return &UserController{Service: *UserService, MailConn: MailConn}
}

func InitializeUserController(api *fiber.Group, DB *gorm.DB, MailConn *gomail.Dialer) {
	userRepo := repository.NewUserRepository(DB)
	userService := service.NewUserService(userRepo.(*repository.UserRepository))
	userController := NewUserController(userService.(*service.UserService), MailConn)
	userController.Route(api)
}

func (controller *UserController) Route(api *fiber.Group) {
	api.Post("/user", controller.Create)
	api.Get("/user", middleware.JWTProtected(), controller.List)
	api.Post("/user/auth", controller.GetNewAccessToken)
	api.Post("/user/reset-password", controller.RequestResetPassword)
}

// CreateUser is a function to insert user to database
// @Summary      Create User
// @Description  Create New User / Register User
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        user  body      model.UserRequest  true  "Register user"
// @Success      200   {object}  model.WebResponse{data=model.UserResponse}
// @Failure      500   {object}  model.WebResponse{data=string}
// @Failure      400   {object}  model.WebResponse{data=string}
// @Router       /v1/user [post]
func (controller *UserController) Create(c *fiber.Ctx) error {
	var request model.UserRequest

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

// GetAllUser is a function to get all user data from database
// @Security ApiKeyAuth
// @Summary      Get All User
// @Description  get all user data from database
// @Tags         user
// @Accept       json
// @Produce      json
// @Success      200   {object}  model.WebResponse{data=[]model.UserResponse}
// @Failure      500   {object}  model.WebResponse{data=string}
// @Failure      400   {object}  model.WebResponse{data=string}
// @Router       /v1/user [get]
func (controller *UserController) List(c *fiber.Ctx) error {
	responses := controller.Service.List()

	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   responses,
	})
}

// AuthUser is a function to authenticate user
// @Summary      Auth User
// @Description  authenticate user / login User
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        auth  body      model.AuthRequest  true  "Auth user"
// @Success      200   {object}  model.WebResponse{data=model.AuthResponse}
// @Failure      500   {object}  model.WebResponse{data=string}
// @Failure      400   {object}  model.WebResponse{data=string}
// @Router       /v1/user/auth [post]
func (controller *UserController) GetNewAccessToken(c *fiber.Ctx) error {
	// Generate a new Access token.
	var request model.AuthRequest
	err := c.BodyParser(&request)
	exception.PanicWhenError(err)

	auth := controller.Service.FindByAuth(request)

	utils.ValidatePassword(request.Password, auth.Password)

	jwtToken := utils.GenerateNewToken(&auth, false)

	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "success",
		Data:   jwtToken,
	})
}

// AuthUser is a function to request reset password user
// @Summary      Request Reset Password User
// @Description  request reset password user
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        data  body      model.ResetPasswordRequest  true  "Request Reset Password User"
// @Success      200   {object}  model.WebResponse{data=string}
// @Failure      500   {object}  model.WebResponse{data=string}
// @Failure      400   {object}  model.WebResponse{data=string}
// @Router       /v1/user/reset-password [post]
func (controller *UserController) RequestResetPassword(c *fiber.Ctx) error {
	var request model.ResetPasswordRequest

	err := c.BodyParser(&request)
	exception.PanicWhenError(err)

	user := controller.Service.FindByEmail(request.Email)

	message := emails.GenerateResetPasswordEmail("http://103.150.60.157:6991/v1/user/reset-password/wkwkwkwkwkkwk")
	message.SetReceiver(user.Email)

	err = utils.SendEmail(controller.MailConn, message)
	exception.PanicWhenError(err)

	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "success",
		Data:   "sukses mengirim email ke " + request.Email,
	})
}

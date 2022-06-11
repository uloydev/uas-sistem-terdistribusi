package controller

import (
	"sister-backend/app/model"
	"sister-backend/app/repository"
	"sister-backend/app/service"
	"sister-backend/exception"
	"strconv"

	"github.com/go-gomail/gomail"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type LiceseController struct {
	Service  service.LicenseService
	MailConn *gomail.Dialer
}

func NewLiceseController(LicenseService *service.LicenseService, MailConn *gomail.Dialer) LiceseController {
	return LiceseController{Service: *LicenseService, MailConn: MailConn}
}

func InitializeLiceseController(api *fiber.Group, DB *gorm.DB, MailConn *gomail.Dialer) {
	licenseRepo := repository.NewLicenseRepository(DB)
	licenseService := service.NewLicenseService(licenseRepo.(*repository.LicenseRepository))
	licenseController := NewLiceseController(licenseService.(*service.LicenseService), MailConn)
	licenseController.Route(api)
}

func (controller *LiceseController) Route(api *fiber.Group) {
	api.Post("/license", controller.Create)
	api.Get("/license", controller.List)
	api.Delete("/license/:id", controller.Delete)
	api.Put("/license/:id", controller.Update)
	api.Get("/license/check", controller.FindByUsernameKey)
}

// CreateLicense is a function to insert License to database
// @Summary      Create License
// @Description  Create New License
// @Tags         license
// @Accept       json
// @Produce      json
// @Param        license  body      model.LicenseRequest  true  "Create License"
// @Success      200   {object}  model.WebResponse{data=model.LicenseResponse}
// @Failure      500   {object}  model.WebResponse{data=string}
// @Failure      400   {object}  model.WebResponse{data=string}
// @Router       /v1/license [post]
func (controller *LiceseController) Create(c *fiber.Ctx) error {
	var request model.LicenseRequest

	err := c.BodyParser(&request)
	exception.PanicWhenError(err)

	response := controller.Service.Create(request)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

// GetAllLicense is a function to get all License data from database
// @Summary      Get All License
// @Description  get all License data from database
// @Tags         license
// @Accept       json
// @Produce      json
// @Success      200   {object}  model.WebResponse{data=[]model.LicenseResponse}
// @Failure      500   {object}  model.WebResponse{data=string}
// @Failure      400   {object}  model.WebResponse{data=string}
// @Router       /v1/license [get]
func (controller *LiceseController) List(c *fiber.Ctx) error {
	responses := controller.Service.List()

	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   responses,
	})
}

// DeleteLicense is a function to delete License from database
// @Summary      Delete License
// @Description  Delete License
// @Tags         license
// @Accept       json
// @Produce      json
// @Param id path int false "int valid" mininum(1)
// @Success      200   {object}  model.WebResponse{data=string}
// @Failure      500   {object}  model.WebResponse{data=string}
// @Failure      400   {object}  model.WebResponse{data=string}
// @Router       /v1/license/{id} [delete]
func (controller *LiceseController) Delete(c *fiber.Ctx) error {
	ID, err := strconv.Atoi(c.Params("id"))
	exception.PanicWhenError(err)

	response := controller.Service.Delete(uint(ID))
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

// UpdateLicense is a function to update License to database
// @Summary      Update License
// @Description  Update License
// @Tags         License
// @Accept       json
// @Produce      json
// @Param 		 id path int false "int valid" mininum(1)
// @Param        license  body      model.LicenseRequest  true  "Update License"
// @Success      200   {object}  model.WebResponse{data=model.LicenseResponse}
// @Failure      500   {object}  model.WebResponse{data=string}
// @Failure      400   {object}  model.WebResponse{data=string}
// @Router       /v1/license/{id} [put]
func (controller *LiceseController) Update(c *fiber.Ctx) error {
	var request model.LicenseRequest

	ID, err := strconv.Atoi(c.Params("id"))
	exception.PanicWhenError(err)

	err = c.BodyParser(&request)
	exception.PanicWhenError(err)

	response := controller.Service.UpdateById(uint(ID), request)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

// CheckLicense is a function to update License to database
// @Summary      Check License
// @Description  Check License
// @Tags         License
// @Accept       json
// @Produce      json
// @Param        license  query      model.CheckLicenseRequest  true  "Check License"
// @Success      200   {object}  model.WebResponse{data=model.LicenseResponse}
// @Failure      500   {object}  model.WebResponse{data=string}
// @Failure      400   {object}  model.WebResponse{data=string}
// @Router       /v1/license/check [get]
func (controller *LiceseController) FindByUsernameKey(c *fiber.Ctx) error {
	var request model.CheckLicenseRequest

	err := c.QueryParser(&request)
	exception.PanicWhenError(err)

	response := controller.Service.FindByUsernameKey(request.Username, request.Key)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

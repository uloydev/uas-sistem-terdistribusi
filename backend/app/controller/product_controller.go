package controller

import (
	"sister-backend/app/model"
	"sister-backend/app/repository"
	"sister-backend/app/service"
	"sister-backend/exception"

	"github.com/go-gomail/gomail"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ProductController struct {
	Service  service.ProductService
	MailConn *gomail.Dialer
}

func NewProductController(ProductService *service.ProductService, MailConn *gomail.Dialer) ProductController {
	return ProductController{Service: *ProductService, MailConn: MailConn}
}

func InitializeProductController(api *fiber.Group, DB *gorm.DB, MailConn *gomail.Dialer) {
	productRepo := repository.NewProductRepository(DB)
	productService := service.NewProductService(productRepo.(*repository.ProductRepository))
	productController := NewProductController(productService.(*service.ProductService), MailConn)
	productController.Route(api)
}

func (controller *ProductController) Route(api *fiber.Group) {
	api.Post("/product", controller.Create)
	api.Get("/product", controller.List)
}

// CreateProduct is a function to insert Product to database
// @Summary      Create Product
// @Description  Create New Product
// @Tags         product
// @Accept       json
// @Produce      json
// @Param        product  body      model.ProductRequest  true  "Create Product"
// @Success      200   {object}  model.WebResponse{data=model.ProductResponse}
// @Failure      500   {object}  model.WebResponse{data=string}
// @Failure      400   {object}  model.WebResponse{data=string}
// @Router       /v1/product [post]
func (controller *ProductController) Create(c *fiber.Ctx) error {
	var request model.ProductRequest

	err := c.BodyParser(&request)
	exception.PanicWhenError(err)

	response := controller.Service.Create(request)
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

// GetAllProduct is a function to get all Product data from database
// @Security ApiKeyAuth
// @Summary      Get All Product
// @Description  get all Product data from database
// @Tags         product
// @Accept       json
// @Produce      json
// @Success      200   {object}  model.WebResponse{data=[]model.ProductResponse}
// @Failure      500   {object}  model.WebResponse{data=string}
// @Failure      400   {object}  model.WebResponse{data=string}
// @Router       /v1/product [get]
func (controller *ProductController) List(c *fiber.Ctx) error {
	responses := controller.Service.List()

	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   responses,
	})
}

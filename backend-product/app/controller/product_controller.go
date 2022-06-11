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
	api.Delete("/product/:id", controller.Delete)
	api.Put("/product/:id", controller.Update)
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

// DeleteProduct is a function to delete Product from database
// @Summary      Delete Product
// @Description  Delete Product
// @Tags         product
// @Accept       json
// @Produce      json
// @Param id path int false "int valid" mininum(1)
// @Success      200   {object}  model.WebResponse{data=string}
// @Failure      500   {object}  model.WebResponse{data=string}
// @Failure      400   {object}  model.WebResponse{data=string}
// @Router       /v1/product/{id} [delete]
func (controller *ProductController) Delete(c *fiber.Ctx) error {
	ID, err := strconv.Atoi(c.Params("id"))
	exception.PanicWhenError(err)

	response := controller.Service.Delete(uint(ID))
	return c.JSON(model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	})
}

// UpdateProduct is a function to update Product to database
// @Security ApiKeyAuth
// @Summary      Update Product
// @Description  Update Product
// @Tags         Product
// @Accept       json
// @Produce      json
// @Param 		 id path int false "int valid" mininum(1)
// @Param        product  body      model.ProductRequest  true  "Update Product"
// @Success      200   {object}  model.WebResponse{data=model.ProductResponse}
// @Failure      500   {object}  model.WebResponse{data=string}
// @Failure      400   {object}  model.WebResponse{data=string}
// @Router       /v1/product/{id} [put]
func (controller *ProductController) Update(c *fiber.Ctx) error {
	var request model.ProductRequest

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

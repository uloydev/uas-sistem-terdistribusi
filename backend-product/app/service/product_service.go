package service

import (
	"sister-backend/app/entity"
	"sister-backend/app/model"
	"sister-backend/app/repository"
	"sister-backend/app/validation"
)

type ProductService struct {
	Repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) BaseService[model.ProductRequest, model.ProductResponse] {
	return &ProductService{
		Repo: repo,
	}
}

func (service *ProductService) Create(request model.ProductRequest) (response model.ProductResponse) {
	validation.ValidateProduct(request)

	product := entity.Product{
		Title:       request.Title,
		Description: request.Description,
		Price:       request.Price,
		Stock:       request.Stock,
		IsMaster:    false,
	}
	product = service.Repo.Insert(product)

	response = model.ProductResponse{
		BasicData: model.BasicData{
			ID:        product.ID,
			CreatedAt: product.CreatedAt,
			UpdatedAt: product.UpdatedAt,
			DeletedAt: product.DeletedAt,
		},
		Title:       product.Title,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		IsMaster:    product.IsMaster,
	}

	return response
}

func (service *ProductService) List() (responses []model.ProductResponse) {
	products := service.Repo.FindAll()

	for _, product := range products {
		responses = append(responses, model.ProductResponse{
			BasicData: model.BasicData{
				ID:        product.ID,
				CreatedAt: product.CreatedAt,
				UpdatedAt: product.UpdatedAt,
				DeletedAt: product.DeletedAt,
			},
			Title:       product.Title,
			Description: product.Description,
			Price:       product.Price,
			Stock:       product.Stock,
			IsMaster:    product.IsMaster,
		})
	}

	return responses
}

func (service *ProductService) FindById(Admin_id uint) (response model.ProductResponse) {
	product := service.Repo.FindById(Admin_id)

	response = model.ProductResponse{
		BasicData: model.BasicData{
			ID:        product.ID,
			CreatedAt: product.CreatedAt,
			UpdatedAt: product.UpdatedAt,
			DeletedAt: product.DeletedAt,
		},
		Title:       product.Title,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		IsMaster:    product.IsMaster,
	}

	return response
}

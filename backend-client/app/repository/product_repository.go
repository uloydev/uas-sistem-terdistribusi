package repository

import (
	"sister-backend/app/entity"
	"sister-backend/exception"

	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) BaseRepository[entity.Product] {
	return &ProductRepository{
		DB: db,
	}
}

func (repo *ProductRepository) Insert(product entity.Product) entity.Product {
	result := repo.DB.Create(&product)
	exception.PanicWhenError(result.Error)
	return product
}

func (repo *ProductRepository) FindAll() (products []entity.Product) {
	result := repo.DB.Find(&products)
	exception.PanicWhenError(result.Error)
	return products
}

func (repo *ProductRepository) FindById(ID uint) (product entity.Product) {
	result := repo.DB.Where("id = ?", ID).First(&product)
	exception.PanicValidationWhenError(result.Error)
	return product
}

func (repo *ProductRepository) Delete(ID uint) {
	result := repo.DB.Delete(&entity.Product{}, ID)
	exception.PanicWhenError(result.Error)
}

func (repo *ProductRepository) UpdateById(product entity.Product) entity.Product {
	result := repo.DB.Model(&product).Updates(map[string]any{
		"title":       product.Title,
		"price":       product.Price,
		"stock":       product.Stock,
		"description": product.Description,
		"is_master":   product.IsMaster,
	}).First(&product)
	exception.PanicValidationWhenError(result.Error)
	return product
}

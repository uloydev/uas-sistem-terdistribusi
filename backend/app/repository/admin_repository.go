package repository

import (
	"sister-backend/app/entity"
	"sister-backend/exception"

	"gorm.io/gorm"
)

type AdminRepository struct {
	DB *gorm.DB
}

func NewAdminRepository(db *gorm.DB) BaseRepository[entity.Admin] {
	return &AdminRepository{
		DB: db,
	}
}

func (repo *AdminRepository) Insert(admin entity.Admin) entity.Admin {
	result := repo.DB.Create(&admin)
	exception.PanicWhenError(result.Error)
	return admin
}

func (repo *AdminRepository) FindAll() (admins []entity.Admin) {
	result := repo.DB.Find(&admins)
	exception.PanicWhenError(result.Error)
	return admins
}

func (repo *AdminRepository) FindById(ID uint) (admin entity.Admin) {
	result := repo.DB.Where("id = ?", ID).First(&admin)
	exception.PanicValidationWhenError(result.Error)
	return admin
}

func (repo *AdminRepository) FindByEmail(email string) (admin entity.Admin) {
	result := repo.DB.Where("email = ?", email).First(&admin)
	exception.PanicValidationWhenError(result.Error)
	return admin
}

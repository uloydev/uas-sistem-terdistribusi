package repository

import (
	"sister-backend/app/entity"
	"sister-backend/exception"

	"gorm.io/gorm"
)

type LicenseRepository struct {
	DB *gorm.DB
}

func NewLicenseRepository(db *gorm.DB) BaseRepository[entity.License] {
	return &LicenseRepository{
		DB: db,
	}
}

func (repo *LicenseRepository) Insert(license entity.License) entity.License {
	result := repo.DB.Create(&license)
	exception.PanicWhenError(result.Error)
	return license
}

func (repo *LicenseRepository) FindAll() (licenses []entity.License) {
	result := repo.DB.Find(&licenses)
	exception.PanicWhenError(result.Error)
	return licenses
}

func (repo *LicenseRepository) FindById(ID uint) (license entity.License) {
	result := repo.DB.Where("id = ?", ID).First(&license)
	exception.PanicValidationWhenError(result.Error)
	return license
}

func (repo *LicenseRepository) Delete(ID uint) {
	result := repo.DB.Delete(&entity.License{}, ID)
	exception.PanicWhenError(result.Error)
}

func (repo *LicenseRepository) UpdateById(license entity.License) entity.License {
	result := repo.DB.Model(&license).Updates(map[string]any{
		"username":  license.Username,
		"key":       license.Key,
		"is_active": license.IsActive,
	}).First(&license)
	exception.PanicValidationWhenError(result.Error)
	return license
}

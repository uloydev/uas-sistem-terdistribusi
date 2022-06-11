package service

import (
	"sister-backend/app/entity"
	"sister-backend/app/model"
	"sister-backend/app/repository"
	"sister-backend/app/validation"
	"strconv"
)

type LicenseService struct {
	Repo *repository.LicenseRepository
}

func NewLicenseService(repo *repository.LicenseRepository) BaseService[model.LicenseRequest, model.LicenseResponse] {
	return &LicenseService{
		Repo: repo,
	}
}

func (service *LicenseService) Create(request model.LicenseRequest) (response model.LicenseResponse) {
	validation.ValidateLicense(request)

	license := entity.License{
		Username: request.Username,
		Key:      request.Key,
		IsActive: request.IsActive,
	}
	license = service.Repo.Insert(license)

	response = model.LicenseResponse{
		BasicData: model.BasicData{
			ID:        license.ID,
			CreatedAt: license.CreatedAt,
			UpdatedAt: license.UpdatedAt,
			DeletedAt: license.DeletedAt,
		},
		Username: license.Username,
		Key:      license.Key,
		IsActive: license.IsActive,
	}

	return response
}

func (service *LicenseService) List() (responses []model.LicenseResponse) {
	licenses := service.Repo.FindAll()

	for _, license := range licenses {
		responses = append(responses, model.LicenseResponse{
			BasicData: model.BasicData{
				ID:        license.ID,
				CreatedAt: license.CreatedAt,
				UpdatedAt: license.UpdatedAt,
				DeletedAt: license.DeletedAt,
			},
			Username: license.Username,
			Key:      license.Key,
			IsActive: license.IsActive,
		})
	}

	return responses
}

func (service *LicenseService) FindById(ID uint) (response model.LicenseResponse) {
	license := service.Repo.FindById(ID)

	response = model.LicenseResponse{
		BasicData: model.BasicData{
			ID:        license.ID,
			CreatedAt: license.CreatedAt,
			UpdatedAt: license.UpdatedAt,
			DeletedAt: license.DeletedAt,
		},
		Username: license.Username,
		Key:      license.Key,
		IsActive: license.IsActive,
	}

	return response
}

func (service *LicenseService) Delete(ID uint) (responses string) {
	service.Repo.Delete(ID)
	return "License with id " + strconv.Itoa(int(ID)) + " deleted."
}

func (service *LicenseService) UpdateById(marketplace_id uint, request model.LicenseRequest) (response model.LicenseResponse) {
	validation.ValidateLicense(request)

	license := entity.License{
		Username: request.Username,
		Key:      request.Key,
		IsActive: request.IsActive,
	}

	license = service.Repo.UpdateById(license)

	response = model.LicenseResponse{
		BasicData: model.BasicData{
			ID:        license.ID,
			CreatedAt: license.CreatedAt,
			UpdatedAt: license.UpdatedAt,
			DeletedAt: license.DeletedAt,
		},
		Username: license.Username,
		Key:      license.Key,
		IsActive: license.IsActive,
	}

	return response
}

func (service *LicenseService) FindByUsernameKey(username string, key string) (response model.LicenseResponse) {
	license := service.Repo.FindByUsernameKey(username, key)

	response = model.LicenseResponse{
		BasicData: model.BasicData{
			ID:        license.ID,
			CreatedAt: license.CreatedAt,
			UpdatedAt: license.UpdatedAt,
			DeletedAt: license.DeletedAt,
		},
		Username: license.Username,
		Key:      license.Key,
		IsActive: license.IsActive,
	}

	return response
}

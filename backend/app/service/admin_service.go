package service

import (
	"sister-backend/app/entity"
	"sister-backend/app/model"
	"sister-backend/app/repository"
	"sister-backend/app/validation"
)

type AdminService struct {
	Repo *repository.AdminRepository
}

func NewAdminService(repo *repository.AdminRepository) BaseService[model.AdminRequest, model.AdminResponse] {
	return &AdminService{
		Repo: repo,
	}
}

func (service *AdminService) Create(request model.AdminRequest) (response model.AdminResponse) {
	validation.ValidateAdmin(request)

	admin := entity.Admin{
		Name:     request.Name,
		Email:    &request.Email,
		Password: &request.Password,
	}
	admin = service.Repo.Insert(admin)

	response = model.AdminResponse{
		BasicData: model.BasicData{
			ID:        admin.ID,
			CreatedAt: admin.CreatedAt,
			UpdatedAt: admin.UpdatedAt,
			DeletedAt: admin.DeletedAt,
		},
		Name:  admin.Name,
		Email: *admin.Email,
	}

	return response
}

func (service *AdminService) List() (responses []model.AdminResponse) {
	admins := service.Repo.FindAll()

	for _, admin := range admins {
		responses = append(responses, model.AdminResponse{
			BasicData: model.BasicData{
				ID:        admin.ID,
				CreatedAt: admin.CreatedAt,
				UpdatedAt: admin.UpdatedAt,
				DeletedAt: admin.DeletedAt,
			},
			Name:  admin.Name,
			Email: *admin.Email,
		})
	}

	return responses
}

func (service *AdminService) FindById(Admin_id uint) (response model.AdminResponse) {
	admin := service.Repo.FindById(Admin_id)

	response = model.AdminResponse{
		BasicData: model.BasicData{
			ID:        admin.ID,
			CreatedAt: admin.CreatedAt,
			UpdatedAt: admin.UpdatedAt,
			DeletedAt: admin.DeletedAt,
		},
		Name:  admin.Name,
		Email: *admin.Email,
	}

	return response
}

func (service *AdminService) FindByEmail(email string) (response model.AdminResponse) {
	admin := service.Repo.FindByEmail(email)

	response = model.AdminResponse{
		BasicData: model.BasicData{
			ID:        admin.ID,
			CreatedAt: admin.CreatedAt,
			UpdatedAt: admin.UpdatedAt,
			DeletedAt: admin.DeletedAt,
		},
		Name:  admin.Name,
		Email: *admin.Email,
	}

	return response
}

func (service *AdminService) FindByAuth(request model.AuthRequest) (response model.AuthResponse) {
	validation.ValidateAuth(request)
	admin := service.Repo.FindByEmail(request.Email)

	response = model.AuthResponse{
		ID:       admin.ID,
		Email:    *admin.Email,
		Password: *admin.Password,
	}

	return response
}

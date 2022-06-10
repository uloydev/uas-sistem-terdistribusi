package service

import (
	"sister-backend/app/entity"
	"sister-backend/app/model"
	"sister-backend/app/repository"
	"sister-backend/app/validation"
)

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) BaseService[model.UserRequest, model.UserResponse] {
	return &UserService{
		Repo: repo,
	}
}

func (service *UserService) Create(request model.UserRequest) (response model.UserResponse) {
	validation.ValidateUser(request)

	user := entity.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
		Birthday: request.Birthday,
		Phone:    request.Phone,
	}

	user = service.Repo.Insert(user)

	response = model.UserResponse{
		BasicData: model.BasicData{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			DeletedAt: user.DeletedAt,
		},
		Name:     user.Name,
		Email:    user.Email,
		Birthday: user.Birthday,
		Phone:    user.Phone,
	}

	return response
}

func (service *UserService) List() (responses []model.UserResponse) {
	users := service.Repo.FindAll()

	for _, user := range users {
		responses = append(responses, model.UserResponse{
			BasicData: model.BasicData{
				ID:        user.ID,
				CreatedAt: user.CreatedAt,
				UpdatedAt: user.UpdatedAt,
				DeletedAt: user.DeletedAt,
			},
			Name:     user.Name,
			Email:    user.Email,
			Birthday: user.Birthday,
			Phone:    user.Phone,
		})
	}

	return responses
}

func (service *UserService) FindById(user_id uint) (response model.UserResponse) {
	user := service.Repo.FindById(user_id)

	response = model.UserResponse{
		BasicData: model.BasicData{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			DeletedAt: user.DeletedAt,
		},
		Name:     user.Name,
		Email:    user.Email,
		Birthday: user.Birthday,
		Phone:    user.Phone,
	}

	return response
}

func (service *UserService) FindByEmail(email string) (response model.UserResponse) {
	user := service.Repo.FindByEmail(email)

	response = model.UserResponse{
		BasicData: model.BasicData{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			DeletedAt: user.DeletedAt,
		},
		Name:     user.Name,
		Email:    user.Email,
		Birthday: user.Birthday,
		Phone:    user.Phone,
	}

	return response
}

func (service *UserService) FindByAuth(request model.AuthRequest) (response model.AuthResponse) {
	validation.ValidateAuth(request)
	user := service.Repo.FindByEmail(request.Email)

	response = model.AuthResponse{
		ID:       user.ID,
		Email:    user.Email,
		Password: user.Password,
	}

	return response
}

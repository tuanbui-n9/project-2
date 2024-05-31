package services

import (
	"context"

	"github.com/tuanbui-n9/project-2/internal/app/models"
	repository "github.com/tuanbui-n9/project-2/internal/app/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	UserRepository *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{
		UserRepository: userRepo,
	}
}

func (us *UserService) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	return us.UserRepository.Create(ctx, user)
}

func (us *UserService) GetUserByID(ctx context.Context, id primitive.ObjectID) (*models.User, error) {
	return us.UserRepository.GetByID(ctx, id)
}

func (us *UserService) UpdateUser(ctx context.Context, id primitive.ObjectID, user *models.User) error {
	return us.UserRepository.Update(ctx, id, user)
}

func (us *UserService) DeleteUser(ctx context.Context, id primitive.ObjectID) error {
	return us.UserRepository.Delete(ctx, id)
}

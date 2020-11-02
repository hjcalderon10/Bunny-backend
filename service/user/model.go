package user

import (
	"context"

	models "github.com/hjcalderon10/bunny-backend/model"
)

type (
	IService interface {
		GetAllUsers(ctx context.Context) ([]models.User, error)
		CreateUser(ctx context.Context, user models.User) (models.User, error)
		ReadUser(ctx context.Context, user *models.User) error
		UpdateUser(ctx context.Context, user models.User) error
		DeleteUser(ctx context.Context, user models.User) error
	}

	IRepo interface {
		GetAllUsers(ctx context.Context) ([]models.User, error)
		CreateUser(ctx context.Context, user models.User) (uint16, error)
		ReadUser(ctx context.Context, userID models.UserID) (models.User, error)
		UpdateUser(ctx context.Context, user models.User) error
		DeleteUser(ctx context.Context, userID models.UserID) error
	}

	service struct {
		repo IRepo
	}
)

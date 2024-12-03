package service

import (
	"context"
	"gochat/internal/models"
)

type UserServ interface {
	Create(ctx context.Context, user models.UserCreate) (int, error)
	Login(ctx context.Context, user models.UserLogin) (int, error)
	Get(ctx context.Context, id int) (*models.User, error)
	Delete(ctx context.Context, id int) error
}

type UserChangeServ interface {
	PWD(ctx context.Context, user models.UserChangePWD) error
	Email(ctx context.Context, user models.UserChangeEmail) error
	Phone(ctx context.Context, user models.UserChangePhone) error
	UserData(ctx context.Context, user models.UserChange) error
}
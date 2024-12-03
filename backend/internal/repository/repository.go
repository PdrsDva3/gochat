package repository

import (
	"context"
	"gochat/internal/models"
)

type UserRepo interface {
	Create(ctx context.Context, user models.UserCreate) (int, error)
	Get(ctx context.Context, id int) (*models.User, error)
	GetPWDbyEmail(ctx context.Context, user string) (int, string, error)
	Delete(ctx context.Context, id int) error
}

type UserChangeRepo interface {
	ChangePWD(ctx context.Context, user models.UserChangePWD) error
	ChangeEmail(ctx context.Context, user models.UserChangeEmail) error
	ChangePhone(ctx context.Context, user models.UserChangePhone) error
	ChangeUserData(ctx context.Context, user models.UserChange) error
}

type FriendRepo interface {
	Get(ctx context.Context, id int) (*models.FriendList, error)
	AddFriend(ctx context.Context, id1 int, id2 int) error
	Delete(ctx context.Context, id1 int, id2 int) error
}

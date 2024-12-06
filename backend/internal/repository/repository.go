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

type ChatRepo interface {
	Create(ctx context.Context, chat models.ChatCreate) (int, error)
	AddUser(ctx context.Context, idUser int, idChat int) error
	List(ctx context.Context, id int) (*models.ChatList, error)
	Get(ctx context.Context, id int) (*models.Chat, error)
	Delete(ctx context.Context, id int) error
}

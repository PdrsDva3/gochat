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

type FriendServ interface {
	Get(ctx context.Context, id int) (*models.FriendList, error)
	AddFriend(ctx context.Context, id1 int, id2 int) error
	Delete(ctx context.Context, id1 int, id2 int) error
}

type ChatServ interface {
	Create(ctx context.Context, chat models.ChatCreate) (int, error)
	AddUser(ctx context.Context, idUser int, idChat int) error
	List(ctx context.Context, id int) (*models.ChatList, error)
	Get(ctx context.Context, id int) (*models.Chat, error)
	Delete(ctx context.Context, id int) error
}

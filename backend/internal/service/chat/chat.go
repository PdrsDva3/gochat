package chat

import (
	"context"
	"fmt"
	"gochat/internal/models"
	"gochat/internal/repository"
	"gochat/internal/service"
	"gochat/pkg/log"
)

type ServChat struct {
	Repo repository.ChatRepo
	log  *log.Logs
}

func InitChatService(chatRepo repository.ChatRepo, log *log.Logs) service.ChatServ {
	return ServChat{Repo: chatRepo, log: log}
}

func (s ServChat) Create(ctx context.Context, chat models.ChatCreate) (int, error) {
	newChat := models.ChatCreate{
		IDCreator:   chat.IDCreator,
		Name:        chat.Name,
		Description: chat.Description,
		IDUsers:     chat.IDUsers,
	}
	id, err := s.Repo.Create(ctx, newChat)
	if err != nil {
		s.log.Error(err.Error())
		return 0, err
	}
	s.log.Info(fmt.Sprintf("create chat : %d", chat.IDCreator))
	return id, nil
}

func (s ServChat) List(ctx context.Context, id int) (*models.ChatList, error) {
	list, err := s.Repo.List(ctx, id)
	if err != nil {
		s.log.Error(err.Error())
		return nil, err
	}
	s.log.Info(fmt.Sprintf("get list chat : %d", id))
	return list, nil
}

func (s ServChat) Get(ctx context.Context, id int) (*models.Chat, error) {
	chat, err := s.Repo.Get(ctx, id)
	if err != nil {
		s.log.Error(err.Error())
		return nil, err
	}
	s.log.Info(fmt.Sprintf("get chat : %d", id))
	return chat, nil
}

func (s ServChat) Delete(ctx context.Context, id int) error {
	err := s.Repo.Delete(ctx, id)
	if err != nil {
		s.log.Error(err.Error())
		return err
	}
	s.log.Info(fmt.Sprintf("delete chat : %d", id))
	return nil
}

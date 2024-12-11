package friend

import (
	"context"
	"fmt"
	"gochat/internal/models"
	"gochat/internal/repository"
	"gochat/internal/service"
	"gochat/pkg/log"
)

type ServFriend struct {
	Repo repository.FriendRepo
	log  *log.Logs
}

func InitFriendService(friendRepo repository.FriendRepo, log *log.Logs) service.FriendServ {
	return ServFriend{Repo: friendRepo, log: log}
}

func (s ServFriend) Get(ctx context.Context, id int) (*models.FriendList, error) {
	friends, err := s.Repo.Get(ctx, id)
	if err != nil {
		s.log.Error(err.Error())
		return nil, err
	}
	s.log.Info(fmt.Sprintf("get friend %v", id))
	return friends, err
}

func (s ServFriend) GetFriendsInfo(ctx context.Context, id int) (*models.FriendListInfo, error) {
	friendsInfo, err := s.Repo.GetFriendsInfo(ctx, id)
	if err != nil {
		s.log.Error(err.Error())
		return nil, err
	}
	s.log.Info(fmt.Sprintf("get friends info %v", friendsInfo))
	return friendsInfo, err
}

func (s ServFriend) AddFriend(ctx context.Context, id1 int, id2 int) error {
	err := s.Repo.AddFriend(ctx, id1, id2)
	if err != nil {
		s.log.Error(err.Error())
		return err
	}
	s.log.Info(fmt.Sprintf("add friend %v-%v", id1, id2))
	return nil
}

func (s ServFriend) Delete(ctx context.Context, id1 int, id2 int) error {
	err := s.Repo.Delete(ctx, id1, id2)
	if err != nil {
		s.log.Error(err.Error())
		return err
	}
	s.log.Info(fmt.Sprintf("delete friend %v-%v", id1, id2))
	return nil
}

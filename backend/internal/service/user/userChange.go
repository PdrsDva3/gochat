package user

import (
	"context"
	"fmt"
	"gochat/internal/models"
	"gochat/internal/repository"
	"gochat/internal/service"
	"gochat/pkg/cerr"
	"gochat/pkg/log"
	"golang.org/x/crypto/bcrypt"
)

type ServUserChange struct {
	Repo repository.UserChangeRepo
	log  *log.Logs
}

func InitUserChangeService(userChangeRepo repository.UserChangeRepo, log *log.Logs) service.UserChangeServ {
	return ServUserChange{Repo: userChangeRepo, log: log}
}

func (s ServUserChange) PWD(ctx context.Context, user models.UserChangePWD) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		s.log.Error(cerr.Err(cerr.Hash, err).Str())
		return cerr.Err(cerr.Hash, err).Error()
	}
	newPWD := models.UserChangePWD{
		ID:       user.ID,
		Password: string(hash),
	}

	err = s.Repo.ChangePWD(ctx, newPWD)
	if err != nil {
		s.log.Error(err.Error())
		return err
	}
	s.log.Info(fmt.Sprintf("change pwd user"))
	return nil
}

func (s ServUserChange) Email(ctx context.Context, user models.UserChangeEmail) error {
	err := s.Repo.ChangeEmail(ctx, user)
	if err != nil {
		s.log.Error(err.Error())
		return err
	}
	return nil
}

func (s ServUserChange) Phone(ctx context.Context, user models.UserChangePhone) error {
	err := s.Repo.ChangePhone(ctx, user)
	if err != nil {
		s.log.Error(err.Error())
		return err
	}
	return nil
}

func (s ServUserChange) UserData(ctx context.Context, user models.UserChange) error {
	err := s.Repo.ChangeUserData(ctx, user)
	if err != nil {
		s.log.Error(err.Error())
		return err
	}
	return nil
}

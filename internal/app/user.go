package app

import (
	"context"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/smtp"

	"github.com/thirteenths/test-bmstu23/internal/app/mapper"
	"github.com/thirteenths/test-bmstu23/internal/domain"
	"github.com/thirteenths/test-bmstu23/internal/domain/requests"
	"github.com/thirteenths/test-bmstu23/internal/domain/responses"
)

type userServiceStorage interface {
	GetUser(ctx context.Context, id int) (domain.User, error)
	GetPassword(ctx context.Context, email string) (domain.User, error)
	UpdatePassword(ctx context.Context, user domain.User) error
}

type UserService struct {
	logger  *log.Logger
	storage userServiceStorage
}

func NewUserService(logger *log.Logger, storage userServiceStorage) *UserService {
	return &UserService{logger: logger, storage: storage}
}

func (s *UserService) GetUser(ctx context.Context, id int) (*responses.GetUser, error) {
	res, err := s.storage.GetUser(ctx, id)
	if err != nil {
		log.WithError(err).Warnf("can't storage.GetUser GetUser")
		return nil, err
	}

	return mapper.MakeResponseUser(res), nil
}

func (s *UserService) CheckPassword(ctx context.Context, user requests.LoginUser) error {
	res, err := s.storage.GetPassword(ctx, user.Email)
	if err != nil {
		log.WithError(err).Warnf("can't storage.GetUser GetUser")
		return err
	}
	if user.Password != res.Hash {
		log.WithError(errors.New("error Password")).Warnf("can't storage.GetUser GetUser")
		return err
	}
	return nil
}

func (s *UserService) UpdatePassword(ctx context.Context, user requests.UpdatePassword) error {
	err := s.storage.UpdatePassword(ctx, *mapper.MakeRequestUpdatePassword(user))
	if err != nil {
		log.WithError(err).Warnf("can't storage.GetUser GetUser")
		return err
	}

	return nil
}

func (s *UserService) SendCode(ctx context.Context, user string) error {
	code := "23456789"

	host := "smtp.gmail.com"
	port := "587"
	login := "t56144938@gmail.com"
	pass := "eoli ehwe uzjt vjwb"

	auth := smtp.PlainAuth("", login, pass, host)
	to := []string{user}
	msg := []byte(fmt.Sprintf("Verify code: %s", code))

	err := smtp.SendMail(fmt.Sprintf("%s:%s", host, port), auth, login, to, msg)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	return nil
}

func (s *UserService) CheckCode(ctx context.Context, code string) error {
	correctCode := "23456789"
	if code != correctCode {
		return errors.New("error code")
	}
	return nil
}

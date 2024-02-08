package mapper

import (
	"github.com/thirteenths/test-bmstu23/internal/domain"
	"github.com/thirteenths/test-bmstu23/internal/domain/requests"
	"github.com/thirteenths/test-bmstu23/internal/domain/responses"
)

func MakeResponseUser(u domain.User) *responses.GetUser {
	return &responses.GetUser{
		ID:       u.ID,
		Email:    u.Email,
		Password: u.Hash,
	}
}

func MakeRequestLogin(l requests.LoginUser) *domain.User {
	return &domain.User{
		Email: l.Email,
		Hash:  l.Password,
	}
}

func MakeRequestUpdatePassword(u requests.UpdatePassword) *domain.User {
	return &domain.User{
		Email: u.Email,
		Hash:  u.Password,
	}
}

func MakeRequestVerifyCode(c requests.VerifyCode) *domain.User {
	return &domain.User{
		Email: c.Email,
		Code:  c.Code,
	}
}

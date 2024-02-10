package storage

import (
	"context"

	"github.com/thirteenths/test-bmstu23/internal/domain"
	"github.com/thirteenths/test-bmstu23/internal/infra/postgres"
)

type Storage interface {
	GetAllEvent(ctx context.Context) ([]domain.Event, error)
	GetEvent(ctx context.Context, id int) (domain.Event, error)
	GetUserEvent(ctx context.Context, idUser int) ([]domain.Event, error)

	CreateEvent(ctx context.Context, event domain.Event) (int, error)
	DeleteEvent(ctx context.Context, id int) error
	UpdateEvent(ctx context.Context, event domain.Event, id int) error

	GetUser(ctx context.Context, id int) (domain.User, error)
	GetPassword(ctx context.Context, email string) (domain.User, error)
	UpdatePassword(ctx context.Context, user domain.User) error
}

type storage struct {
	postgres postgres.Postgres
}

func NewStorage(postgres postgres.Postgres) *storage {
	return &storage{postgres: postgres}
}

func (s *storage) GetAllEvent(ctx context.Context) ([]domain.Event, error) {
	return s.postgres.GetAllEvent()
}

func (s *storage) GetEvent(ctx context.Context, id int) (domain.Event, error) {
	return s.postgres.GetEvent(id)
}

func (s *storage) GetUserEvent(ctx context.Context, idUser int) ([]domain.Event, error) {
	return s.postgres.GetUserEvent(idUser)
}

func (s *storage) CreateEvent(ctx context.Context, event domain.Event) (int, error) {
	return s.postgres.CreateEvent(event)
}

func (s *storage) DeleteEvent(ctx context.Context, id int) error {
	return s.postgres.DeleteEvent(id)
}

func (s *storage) UpdateEvent(ctx context.Context, event domain.Event, id int) error {
	return s.postgres.UpdateEvent(event, id)
}

func (s *storage) GetUser(ctx context.Context, id int) (domain.User, error) {
	return s.postgres.GetUser(id)
}
func (s *storage) GetPassword(ctx context.Context, email string) (domain.User, error) {
	return s.postgres.GetPassword(email)
}

func (s *storage) UpdatePassword(ctx context.Context, user domain.User) error {
	return s.postgres.UpdatePassword(user)
}

package app

import (
	"context"
	log "github.com/sirupsen/logrus"

	"github.com/thirteenths/test-bmstu23/internal/app/mapper"
	"github.com/thirteenths/test-bmstu23/internal/domain"
	"github.com/thirteenths/test-bmstu23/internal/domain/requests"
	"github.com/thirteenths/test-bmstu23/internal/domain/responses"
)

type eventServiceStorage interface {
	GetAllEvent(ctx context.Context) ([]domain.Event, error)
	GetEvent(ctx context.Context, id int) (domain.Event, error)
	GetUserEvent(ctx context.Context, idUser int) ([]domain.Event, error)

	CreateEvent(ctx context.Context, event domain.Event) (int, error)
	DeleteEvent(ctx context.Context, id int) error
	UpdateEvent(ctx context.Context, event domain.Event, id int) error
}

type EventService struct {
	logger  *log.Logger
	storage eventServiceStorage
}

func NewEventService(logger *log.Logger, storage eventServiceStorage) *EventService {
	return &EventService{logger: logger, storage: storage}
}

func (s *EventService) GetAllEvent(ctx context.Context) (*responses.GetAllEvent, error) {
	res, err := s.storage.GetAllEvent(ctx)
	if err != nil {
		log.WithError(err).Warnf("can't storage.GetAllEvent GetAllEvent")
		return nil, err
	}
	return mapper.MakeResponseAllEvent(res), nil
}

func (s *EventService) GetEvent(ctx context.Context, id int) (*responses.GetEvent, error) {
	res, err := s.storage.GetEvent(ctx, id)
	if err != nil {
		log.WithError(err).Warnf("can't storage.GetEvent GetEvent")
		return nil, err
	}

	return mapper.MakeResponseEvent(res), nil
}

func (s *EventService) GetUserEvent(ctx context.Context, idUser int) (*responses.GetAllEvent, error) {
	res, err := s.storage.GetUserEvent(ctx, idUser)
	if err != nil {
		log.WithError(err).Warnf("can't storage.GetAllEvent GetAllEvent")
		return nil, err
	}
	return mapper.MakeResponseAllEvent(res), nil
}

func (s *EventService) CreateEvent(ctx context.Context, event requests.CreateEvent) (int, error) {
	id, err := s.storage.CreateEvent(ctx, mapper.MakeRequestCreateEvent(event))
	if err != nil {
		log.WithError(err).Warnf("can't storage.CreateEvent CreateEvent")
		return 0, err
	}

	return id, nil
}

func (s *EventService) DeleteEvent(ctx context.Context, id int) error {
	err := s.storage.DeleteEvent(ctx, id)
	if err != nil {
		log.WithError(err).Warnf("can't storage.DeleteEvent DeleteEvent")
		return err
	}

	return nil
}

func (s *EventService) UpdateEvent(ctx context.Context, event requests.UpdateEvent, id int) error {
	err := s.storage.UpdateEvent(ctx, mapper.RequestUpdateEvent(event), id)
	if err != nil {
		log.WithError(err).Warnf("can't storage.UpdateEvent UpdateEvent")
		return err
	}

	return nil
}

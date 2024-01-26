package mapper

import (
	"github.com/thirteenths/test-bmstu23/internal/domain"
	"github.com/thirteenths/test-bmstu23/internal/domain/requests"
	"github.com/thirteenths/test-bmstu23/internal/domain/responses"
)

func MakeResponseAllEvent(e []domain.Event) *responses.GetAllEvent {
	var events []responses.Event
	for _, i := range e {
		events = append(events,
			responses.Event{
				ID:          i.ID,
				Name:        i.Name,
				Description: i.Description,
				Date:        i.Date,
			})
	}
	return &responses.GetAllEvent{Event: events}
}

func MakeResponseEvent(e domain.Event) *responses.GetEvent {
	return &responses.GetEvent{
		ID:          e.ID,
		Name:        e.Name,
		Description: e.Description,
		Date:        e.Date,
	}
}

func MakeRequestCreateEvent(e requests.CreateEvent) domain.Event {
	return domain.Event{
		Name:        e.Name,
		Description: e.Description,
		Date:        e.Date,
	}
}

func RequestUpdateEvent(e requests.UpdateEvent) domain.Event {
	return domain.Event{
		Name:        e.Name,
		Description: e.Description,
		Date:        e.Date,
	}
}

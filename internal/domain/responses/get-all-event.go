package responses

import "time"

type GetAllEvent struct {
	Event []Event `json:"event"`
}

type Event struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

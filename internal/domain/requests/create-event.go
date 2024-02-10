package requests

import "time"

type CreateEvent struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

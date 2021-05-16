package models

import "time"

type Reminder struct {
	ID         int           `json:"id"`
	Title      string        `json:"title"`
	Message    string        `json:"message"`
	Duration   time.Duration `json:"duration"`
	CreatedAt  time.Time     `json:"createdAt"`
	ModifiedAt time.Time     `json:"modifiedAt"`
}

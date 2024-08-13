package models

import "time"

type ApiCollection struct {
	ID          uint64
	Title       string
	Description string
	Collection  string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

package models

import "time"

type TestRecord struct {
	ID         uint64
	TestDataID uint64
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}

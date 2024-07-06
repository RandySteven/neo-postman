package models

import (
	"encoding/json"
	"time"
)

type JiraIssue struct {
	ID        uint64          `json:"id"`
	Link      string          `json:"link"`
	Request   json.RawMessage `json:"request"`
	Response  json.RawMessage `json:"response"`
	CreatedAt time.Time       `json:"createdAt"`
	CreatedBy string          `json:"createdBy"`
	UpdatedAt time.Time       `json:"updatedAt"`
	UpdatedBy string          `json:"updatedBy"`
	DeletedAt *time.Time      `json:"deletedAt"`
}

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
	UpdatedAt time.Time       `json:"updatedAt"`
}

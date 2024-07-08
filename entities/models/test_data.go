package models

import (
	"encoding/json"
	"github.com/RandySteven/neo-postman/enums"
	"time"
)

type TestData struct {
	ID                   uint64
	Method               string
	URI                  string
	Description          string
	RequestHeader        json.RawMessage
	RequestBody          json.RawMessage
	ExpectedResponseCode int
	ExpectedResponse     json.RawMessage
	ActualResponseCode   int
	ActualResponse       json.RawMessage
	ResultStatus         enums.ResultStatus
	IsSaved              bool
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DeletedAt            *time.Time
}

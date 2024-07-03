package models

import (
	"go-api-test/enums"
	"time"
)

type TestData struct {
	ID                   uint64
	Method               string
	URI                  string
	Description          string
	RequestHeader        map[string]interface{}
	RequestBody          map[string]interface{}
	ExpectedResponseCode int
	ExpectedResponse     map[string]interface{}
	ActualResponseCode   int
	ActualResponse       map[string]interface{}
	ResultStatus         enums.ResultStatus
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DeletedBy            *time.Time
}

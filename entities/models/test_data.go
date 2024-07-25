package models

import (
	"encoding/json"
	"github.com/RandySteven/neo-postman/enums"
	"time"
)

type TestData struct {
	ID                   uint64
	Method               string
	Host                 string
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
	ResponseTime         time.Duration

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (t *TestData) JsonRequest() (requestHeaderStr, requestBodyStr, expectedResponseStr, actualResponseStr string) {
	requestHeaderByte, _ := t.RequestHeader.MarshalJSON()
	requestHeaderStr = string(requestHeaderByte)
	requestBodyByte, _ := t.RequestBody.MarshalJSON()
	requestBodyStr = string(requestBodyByte)
	expectedResponseByte, _ := t.ExpectedResponse.MarshalJSON()
	expectedResponseStr = string(expectedResponseByte)
	actualResponseByte, _ := t.ActualResponse.MarshalJSON()
	actualResponseStr = string(actualResponseByte)
	return
}

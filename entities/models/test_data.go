package models

import (
	"encoding/json"
	"github.com/RandySteven/neo-postman/enums"
	"time"
)

type TestData struct {
	ID                   uint64             `json:"id"`
	Method               string             `json:"method"`
	Host                 string             `json:"host"`
	URI                  string             `json:"uri"`
	Description          string             `json:"description"`
	RequestHeader        json.RawMessage    `json:"request_header"`
	RequestBody          json.RawMessage    `json:"request_body"`
	ExpectedResponseCode int                `json:"expected_response_code"`
	ExpectedResponse     json.RawMessage    `json:"expected_response"`
	ActualResponseCode   int                `json:"actual_response_code"`
	ActualResponse       json.RawMessage    `json:"actual_response"`
	ResultStatus         enums.ResultStatus `json:"result_status"`
	IsSaved              bool               `json:"is_saved"`
	ResponseTime         time.Duration      `json:"response_time"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
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

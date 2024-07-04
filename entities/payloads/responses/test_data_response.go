package responses

import (
	"encoding/json"
	"time"
)

type (
	TestDataResponse struct {
		ID                   uint64          `json:"id"`
		ResultStatus         string          `json:"result_status"`
		ExpectedResponseCode int             `json:"expected_response_code,omitempty"`
		ActualResponseCode   int             `json:"actual_response_code,omitempty"`
		ExpectedResponseBody json.RawMessage `json:"expected_response_body,omitempty"`
		ActualResponseBody   json.RawMessage `json:"actual_response_body,omitempty"`
	}

	TestRecordList struct {
		ID           uint64    `json:"id"`
		Description  string    `json:"description"`
		ResultStatus string    `json:"result_status"`
		CreatedAt    time.Time `json:"created_at"`
	}

	TestRecordDetail struct {
		ID               uint64                   `json:"id"`
		Endpoint         string                   `json:"endpoint"`
		Method           string                   `json:"method"`
		Description      string                   `json:"description"`
		ExpectedResponse TestDataExpectedResponse `json:"expected_response"`
		ActualResponse   TestDataActualResponse   `json:"actual_response"`
		ResultStatus     string                   `json:"result_status"`
		CreatedAt        time.Time                `json:"created_at"`
	}

	TestDataExpectedResponse struct {
		ExpectedResponseCode int             `json:"response_code"`
		ExpectedResponse     json.RawMessage `json:"response_body"`
	}

	TestDataActualResponse struct {
		ActualResponseCode int             `json:"response_code"`
		ActualResponse     json.RawMessage `json:"response_body"`
	}
)

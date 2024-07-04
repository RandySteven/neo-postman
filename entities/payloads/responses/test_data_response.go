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
)

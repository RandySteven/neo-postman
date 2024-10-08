package responses

import (
	"encoding/json"
	"time"
)

type (
	TestRecordCreateResponse struct {
		ID uint64 `json:"id,omitempty"`
	}

	TestRecordDetailResponse struct {
		ID       uint64 `json:"id"`
		TestData struct {
			ID            uint64          `json:"id"`
			RequestHeader json.RawMessage `json:"request_header"`
			RequestBody   json.RawMessage `json:"request_body"`
			ResultStatus  string          `json:"result_status"`
			ResponseCode  int             `json:"response_code"`
			Links         struct {
				Detail string `json:"detail"`
			} `json:"links"`
		} `json:"test_data"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	TestRecordListResponse struct {
		ID          uint64 `json:"id"`
		Endpoint    string `json:"endpoint"`
		Description string `json:"description"`
		Links       struct {
			Detail string `json:"detail"`
		} `json:"links"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)

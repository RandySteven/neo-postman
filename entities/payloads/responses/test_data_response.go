package responses

type TestDataResponse struct {
	ID                   uint64                 `json:"id"`
	ResultStatus         string                 `json:"result_status"`
	ExpectedResponseCode int                    `json:"expected_response_code,omitempty"`
	ActualResponseCode   int                    `json:"actual_response_code,omitempty"`
	ExpectedResponseBody map[string]interface{} `json:"expected_response_body,omitempty"`
	ActualResponseBody   map[string]interface{} `json:"actual_response_body,omitempty"`
}

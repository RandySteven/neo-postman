package requests

type TestDataRequest struct {
	Method               string                 `json:"method" validate:"required"`
	Path                 string                 `json:"path" validate:"required"`
	Description          string                 `json:"description"`
	RequestHeader        map[string]interface{} `json:"request_header" validate:"required"`
	RequestBody          map[string]interface{} `json:"request_body" validate:"required"`
	ExpectedResponseCode int                    `json:"expected_response_code" validate:"required"`
	ExpectedResponse     map[string]interface{} `json:"expected_response" validate:"required"`
}

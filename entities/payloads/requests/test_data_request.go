package requests

import "encoding/json"

type TestDataRequest struct {
	Method               string          `json:"method" validate:"required"`
	URLKey               string          `json:"url_key" validate:"required"`
	Path                 string          `json:"path" validate:"required"`
	Description          string          `json:"description"`
	RequestHeader        json.RawMessage `json:"request_header" validate:"required"`
	RequestBody          json.RawMessage `json:"request_body" validate:"required"`
	ExpectedResponseCode int             `json:"expected_response_code" validate:"required"`
	ExpectedResponse     json.RawMessage `json:"expected_response" validate:"required"`
}

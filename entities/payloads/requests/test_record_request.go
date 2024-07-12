package requests

type TestRecordRequest struct {
	TestDataID  uint64   `json:"test_data_id;omitempty"`
	TestDataIDs []uint64 `json:"test_data_ids;omitempty"`
}

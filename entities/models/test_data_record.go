package models

type TestDataRecord struct {
	TestDataID   uint64     `json:"test_data_id"`
	TestRecordID uint64     `json:"test_record_id"`
	TestData     TestData   `json:"test_data"`
	TestRecord   TestRecord `json:"test_record"`
}

package queries

const (
	InsertTestRecord GoQuery = `
		INSERT INTO test_records (test_data_id) VALUES ($1)
		RETURNING id
	`

	SelectAllTestRecords GoQuery = `
		SELECT 
		    id, test_data_id, created_at, updated_at, deleted_at 
		FROM test_records 
		WHERE 
		    deleted_at IS NULL
	`

	SelectTestRecordById GoQuery = `
		SELECT id, test_data_id, created_at, updated_at, deleted_at
		FROM test_records
		WHERE id = $1
	`
)

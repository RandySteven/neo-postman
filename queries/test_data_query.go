package queries

const (
	InsertTestData GoQuery = `
		INSERT INTO test_datas 
		    (method, host, uri, description, request_header, request_body, expected_response_code, expected_response, actual_response_code, actual_response, result_status, response_time)
		VALUES 
		    ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		RETURNING id
	`

	SelectTestData GoQuery = `
		SELECT 
				id, method, host, uri, description, request_header, 
				request_body, expected_response_code, expected_response, 
				actual_response_code, actual_response, result_status, is_saved, response_time, created_at, updated_at, deleted_at
			FROM test_datas
		WHERE id = $1
	`

	SelectAllTestData GoQuery = `
		SELECT 
				id, method, host, uri, description, request_header, 
				request_body, expected_response_code, expected_response, 
				actual_response_code, actual_response, result_status, is_saved, response_time, created_at, updated_at, deleted_at
			FROM test_datas
	`

	UpdateTestData GoQuery = ` 
		UPDATE test_datas
		SET 
		    method = $1,
		    host = $2,
		    uri = $3,
		    description = $4,
		    request_header = $5,
		    request_body = $6,
		    expected_response_code = $7,
		    expected_response = $8,
		    actual_response_code = $9,
		    actual_response = $10,
		    result_status = $11,
		    is_saved = $12,
		    updated_at = NOW()
		WHERE id = $13
	`

	DeleteTestUnsavedDatas GoQuery = `
		DELETE FROM test_datas
		WHERE is_saved = false
	`
)

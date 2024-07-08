package queries

const (
	InsertTestData GoQuery = `
		INSERT INTO test_datas 
		    (method, uri, description, request_header, request_body, expected_response_code, expected_response, actual_response_code, actual_response, result_status)
		VALUES 
		    ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING id
	`

	SelectTestData GoQuery = `
		SELECT 
				id, method, uri, description, request_header, 
				request_body, expected_response_code, expected_response, 
				actual_response_code, actual_response, result_status, is_saved, created_at, updated_at, deleted_at
			FROM test_datas
		WHERE id = $1
	`

	SelectAllTestData GoQuery = `
		SELECT 
				id, method, uri, description, request_header, 
				request_body, expected_response_code, expected_response, 
				actual_response_code, actual_response, result_status, is_saved, created_at, updated_at, deleted_at
			FROM test_datas
	`

	UpdateTestData GoQuery = ` 
		UPDATE test_datas
		SET 
		    method = $1,
		    uri = $2,
		    description = $3,
		    request_header = $4,
		    request_body = $5,
		    expected_response_code = $6,
		    expected_response = $7,
		    actual_response_code = $8,
		    actual_response = $9,
		    result_status = $10,
		    is_saved = $11,
		    updated_at = NOW(),
		WHERE id = $12
	`
)

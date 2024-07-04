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
			actual_response_code, actual_response, result_status, created_at, updated_at, deleted_at
		FROM test_datas
	`
)

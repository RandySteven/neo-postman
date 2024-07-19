package queries

const (
	CreateTestDataTable MigrationQuery = `
		CREATE TABLE IF NOT EXISTS test_datas (
		    id BIGSERIAL PRIMARY KEY,
		    method VARCHAR NOT NULL,
		    host VARCHAR NOT NULL,
		    uri VARCHAR NOT NULL,
		    description VARCHAR NOT NULL,
		    request_header JSONB NOT NULL,
		    request_body JSONB NOT NULL,
		    expected_response_code INTEGER NOT NULL,
		    expected_response JSONB NOT NULL,
		    actual_response_code INTEGER NOT NULL,
		    actual_response JSONB NOT NULL,
		    result_status INTEGER NOT NULL,
		    is_saved BOOLEAN NOT NULL DEFAULT FALSE,
		    response_time NUMERIC DEFAULT 0,
		    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
		    deleted_at TIMESTAMP DEFAULT NULL
		)
	`

	CreateJiraIssueTable MigrationQuery = `
		CREATE TABLE IF NOT EXISTS jira_issues (
		    id BIGSERIAL PRIMARY KEY,
		    request JSONB NOT NULL,
		    response JSONB NOT NULL,
		    link VARCHAR NOT NULL,
		    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
		    created_by VARCHAR NOT NULL,
		    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
		    updated_by VARCHAR NOT NULL,
		    deleted_at TIMESTAMP DEFAULT NULL
		)
	`

	CreateTestRecordTable MigrationQuery = ` 
		CREATE TABLE IF NOT EXISTS test_records (
			id BIGSERIAL PRIMARY KEY,
			test_data_id BIGINT NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
			deleted_at TIMESTAMP DEFAULT NULL,
			FOREIGN KEY (test_data_id) REFERENCES test_datas(id)
		)
	`
)

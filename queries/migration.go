package queries

const (
	CreateTestDataTable MigrationQuery = `
		CREATE TABLE IF NOT EXISTS test_datas (
		    id BIGSERIAL PRIMARY KEY,
		    method VARCHAR NOT NULL,
		    uri VARCHAR NOT NULL,
		    description VARCHAR NOT NULL,
		    request_header JSONB NOT NULL,
		    request_body JSONB NOT NULL,
		    expected_response_code INTEGER NOT NULL,
		    expected_response JSONB NOT NULL,
		    actual_response_code INTEGER NOT NULL,
		    actual_response JSONB NOT NULL,
		    result_status INTEGER NOT NULL,
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
)

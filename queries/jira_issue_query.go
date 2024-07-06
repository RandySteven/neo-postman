package queries

const (
	InsertJiraIssue GoQuery = `
		INSERT INTO jira_issues (request, response, link, created_by, updated_by)
		VALUES 
		    ($1, $2, $3, $4, $5)
		RETURNING id
    `

	SelectAllJiraIssues GoQuery = `
		SELECT 
		    id, request, response, link, created_at, created_by, updated_at, updated_by, deleted_at 
		FROM jira_issues
	`
)

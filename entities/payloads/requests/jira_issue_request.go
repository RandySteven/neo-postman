package requests

import "github.com/RandySteven/neo-postman/enums"

type (
	CreateJiraIssueRequest struct {
		Project struct {
			Key string `json:"key"`
		} `json:"project"`
		Summary     string `json:"summary"`
		Description string `json:"description"`
		IssueType   struct {
			Name enums.IssueType `json:"name"`
		} `json:"issueType"`
	}
)

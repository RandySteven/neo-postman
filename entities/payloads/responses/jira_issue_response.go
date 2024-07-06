package responses

import "time"

type (
	CreateJiraIssueResponse struct {
		ID   uint64 `json:"id"`
		Link string `json:"link"`
	}

	JiraIssueListResponse struct {
		ID        uint64    `json:"id"`
		Link      string    `json:"link"`
		CreatedAt time.Time `json:"created_at"`
	}
)

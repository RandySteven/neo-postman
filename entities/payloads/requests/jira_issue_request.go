package requests

type (
	CreateJiraIssueRequest struct {
		Project struct {
			Key string `json:"key"`
		} `json:"project"`
		Summary   string `json:"summary"`
		IssueType struct {
			Name string `json:"name"`
		} `json:"issueType"`
	}
)

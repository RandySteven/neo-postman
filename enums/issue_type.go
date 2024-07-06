package enums

type IssueType string

const (
	Bug   IssueType = "bug"
	Task  IssueType = "task"
	Story IssueType = "story"
)

func (e IssueType) ToString() string {
	return string(e)
}

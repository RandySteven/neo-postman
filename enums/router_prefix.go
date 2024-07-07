package enums

type RouterPrefix string

const (
	TestDataPrefix  RouterPrefix = "/testdata"
	JiraIssuePrefix RouterPrefix = "/jira"
)

func (r RouterPrefix) ToString() string {
	return string(r)
}

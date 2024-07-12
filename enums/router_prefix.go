package enums

type RouterPrefix string

const (
	TestDataPrefix   RouterPrefix = "/testdata"
	JiraIssuePrefix  RouterPrefix = "/jira"
	TestRecordPrefix RouterPrefix = "/testrecord"
)

func (r RouterPrefix) ToString() string {
	return string(r)
}

package enums

type RouterPrefix string

const (
	DevPrefix        RouterPrefix = "/dev"
	TestDataPrefix   RouterPrefix = "/testdata"
	JiraIssuePrefix  RouterPrefix = "/jira"
	TestRecordPrefix RouterPrefix = "/testrecord"
)

func (r RouterPrefix) ToString() string {
	return string(r)
}

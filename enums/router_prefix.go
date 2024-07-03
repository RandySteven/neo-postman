package enums

type RouterPrefix string

const (
	TestDataPrefix RouterPrefix = "/testdata"
)

func (r RouterPrefix) ToString() string {
	return string(r)
}

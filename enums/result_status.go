package enums

type ResultStatus int

const (
	Expected ResultStatus = iota + 1
	Unexpected
	Error
)

func (s ResultStatus) ToString() string {
	switch s {
	case Expected:
		return "expected"
	case Unexpected:
		return "unexpected"
	case Error:
		return "error"
	}
	return "unknown"
}

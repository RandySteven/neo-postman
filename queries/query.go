package queries

const (
	AND = `AND`
	OR  = `OR`
	IN  = `IN`
)

type (
	MigrationQuery string
	DropQuery      string
	GoQuery        string

	Param struct {
		Operator string
		Value    interface{}
	}

	QueryParam struct {
		Params map[string]Param
	}
)

func (q MigrationQuery) ToString() string {
	return string(q)
}

func (q DropQuery) ToString() string {
	return string(q)
}

func (q GoQuery) ToString() string {
	return string(q)
}

package queries

const (
	AND   = `AND`
	OR    = `OR`
	IN    = `IN`
	NOT   = `NOT`
	EQUAL = `=`
)

type (
	MigrationQuery string
	DropQuery      string
	GoQuery        string

	Param struct {
		Operator          string
		ConnectorOperator string
		Value             interface{}
	}

	QueryParam struct {
		Params map[string]Param
		Page   int
		Limit  int
	}
)

func (p *QueryParam) ToString() string {
	//for k, v := range p.Params {
	//	query := fmt.Sprintf("%s %s %s", v.Operator, k, v.ConnectorOperator)
	//}
	return ""
}

func (q MigrationQuery) ToString() string {
	return string(q)
}

func (q DropQuery) ToString() string {
	return string(q)
}

func (q GoQuery) ToString() string {
	return string(q)
}

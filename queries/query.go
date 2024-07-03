package queries

type (
	MigrationQuery string
	DropQuery      string
	GoQuery        string
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

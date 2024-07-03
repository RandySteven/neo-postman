package enums

type ContextKey int

const (
	RequestID ContextKey = iota + 1
	UserID
)

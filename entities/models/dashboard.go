package models

type (
	ExpectedResultCount struct {
		Expected   uint64 `json:"expected"`
		Unexpected uint64 `json:"unexpected"`
	}
)

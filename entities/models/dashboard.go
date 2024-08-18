package models

type (
	ExpectedResultCount struct {
		Expected   uint64 `json:"expected"`
		Unexpected uint64 `json:"unexpected"`
	}

	AvgResponseTimePerApi struct {
		Uri     string  `json:"uri"`
		AvgTime float64 `json:"avg_time"`
	}
)

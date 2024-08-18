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

	CountApiMethod struct {
		Post   uint64 `json:"post"`
		Get    uint64 `json:"get"`
		Patch  uint64 `json:"patch"`
		Put    uint64 `json:"put"`
		Delete uint64 `json:"delete"`
	}
)

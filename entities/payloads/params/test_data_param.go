package params

import (
	"github.com/RandySteven/neo-postman/enums"
	"net/url"
	"strconv"
)

type TestDataParam struct {
	IsSave       bool
	ResultStatus enums.ResultStatus
}

func NewTestParam(r url.Values) *TestDataParam {
	isSaved := r.Get("is_saved")
	resultStatus := r.Get("result_status")
	resultStatusInt, _ := strconv.Atoi(resultStatus)
	return &TestDataParam{
		IsSave:       isSaved == "true",
		ResultStatus: enums.ResultStatus(resultStatusInt),
	}
}

func TestDataParamValidation() string {
	return ""
}

package handlers

import (
	"context"
	"github.com/RandySteven/neo-postman/enums"
	handlers_interfaces "github.com/RandySteven/neo-postman/interfaces/handlers"
	usecases_interfaces "github.com/RandySteven/neo-postman/interfaces/usecases"
	"github.com/RandySteven/neo-postman/utils"
	"github.com/google/uuid"
	"net/http"
)

type DashboardHandler struct {
	usecase usecases_interfaces.DashboardUsecase
}

func (d *DashboardHandler) GetActiveTools(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = context.WithValue(r.Context(), enums.RequestID, uuid.NewString())
		dataKey = `result`
	)
	activeServices := []struct {
		Key   string `json:"key"`
		Value any    `json:"value"`
		Image string `json:"image,omitempty"`
	}{
		{
			Key:   "Redis",
			Value: ctx.Value(enums.ActiveRedis),
			Image: "",
		},
		{
			Key:   "Postgres",
			Value: ctx.Value(enums.ActivePostgres),
			Image: "",
		},
		{
			Key:   "Elastic",
			Value: ctx.Value(enums.ActiveElastic),
			Image: "",
		},
	}
	utils.ResponseHandler(w, http.StatusOK, `success get response`, &dataKey, activeServices, nil)
}

func (d *DashboardHandler) GetCountMethod(w http.ResponseWriter, r *http.Request) {
	var (
		rID     = uuid.NewString()
		ctx     = context.WithValue(r.Context(), enums.RequestID, rID)
		dataKey = `results`
	)
	result, customErr := d.usecase.GetMethodCount(ctx)
	if customErr != nil {
		utils.ResponseHandler(w, customErr.ErrCode(), `failed to get response`, nil, nil, customErr)
		return
	}
	utils.ResponseHandler(w, http.StatusOK, `success get response`, &dataKey, result, nil)
}

func (d *DashboardHandler) GetAvgResponseTimePerAPIs(w http.ResponseWriter, r *http.Request) {
	var (
		rID     = uuid.NewString()
		ctx     = context.WithValue(r.Context(), enums.RequestID, rID)
		dataKey = `results`
	)
	result, customErr := d.usecase.GetAvgResponseTimePerAPIs(ctx)
	if customErr != nil {
		utils.ResponseHandler(w, customErr.ErrCode(), `failed to get response`, nil, nil, customErr)
		return
	}
	utils.ResponseHandler(w, http.StatusOK, `success get response`, &dataKey, result, nil)
}

func (d *DashboardHandler) GetExpectedUnexpectedResult(w http.ResponseWriter, r *http.Request) {
	var (
		rID     = uuid.NewString()
		ctx     = context.WithValue(r.Context(), enums.RequestID, rID)
		dataKey = `results`
	)
	result, customErr := d.usecase.GetExpectedUnexpectedResult(ctx)
	if customErr != nil {
		utils.ResponseHandler(w, customErr.ErrCode(), `failed to get response`, nil, nil, customErr)
		return
	}
	utils.ResponseHandler(w, http.StatusOK, `success get response`, &dataKey, result, nil)
}

var _ handlers_interfaces.DashboardHandler = &DashboardHandler{}

func NewDashboardHandler(usecase usecases_interfaces.DashboardUsecase) *DashboardHandler {
	return &DashboardHandler{
		usecase: usecase,
	}
}

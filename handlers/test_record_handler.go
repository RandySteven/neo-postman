package handlers

import (
	"context"
	"github.com/RandySteven/neo-postman/entities/payloads/requests"
	"github.com/RandySteven/neo-postman/enums"
	handlers_interfaces "github.com/RandySteven/neo-postman/interfaces/handlers"
	usecases_interfaces "github.com/RandySteven/neo-postman/interfaces/usecases"
	"github.com/RandySteven/neo-postman/utils"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type TestRecordHandler struct {
	testRecordUsecase usecases_interfaces.TestRecordUseCase
}

func (t *TestRecordHandler) CreateTestRecord(w http.ResponseWriter, r *http.Request) {
	var (
		rID     = uuid.NewString()
		ctx     = context.WithValue(r.Context(), enums.RequestID, rID)
		request = &requests.TestRecordRequest{}
	)
	if err := utils.BindRequest(r, &request); err != nil {
		utils.ResponseHandler(w, http.StatusBadRequest, `request invalid`, nil, nil, err)
		return
	}
	customErr := t.testRecordUsecase.CreateTestRecord(ctx, request)
	if customErr != nil {
		utils.ResponseHandler(w, customErr.ErrCode(), `request invalid`, nil, nil, customErr)
		return
	}
	utils.ResponseHandler(w, http.StatusCreated, `success create test record`, nil, nil, nil)
}

func (t *TestRecordHandler) GetAllTestRecords(w http.ResponseWriter, r *http.Request) {
	var (
		rID     = uuid.NewString()
		ctx     = context.WithValue(r.Context(), enums.RequestID, rID)
		dataKey = `records`
	)
	result, customErr := t.testRecordUsecase.GetAllTestRecords(ctx)
	if customErr != nil {
		utils.ResponseHandler(w, customErr.ErrCode(), `request invalid`, nil, nil, customErr)
		return
	}
	utils.ResponseHandler(w, http.StatusOK, `success get all records`, &dataKey, result, nil)
}

func (t *TestRecordHandler) GetTestRecordDetail(w http.ResponseWriter, r *http.Request) {
	var (
		rID     = uuid.NewString()
		ctx     = context.WithValue(r.Context(), enums.RequestID, rID)
		dataKey = `record`
		params  = mux.Vars(r)
	)
	id := params[`id`]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		utils.ResponseHandler(w, http.StatusBadRequest, `request invalid`, nil, nil, err)
		return
	}
	result, customErr := t.testRecordUsecase.GetTestRecordDetail(ctx, uint64(idInt))
	if customErr != nil {
		utils.ResponseHandler(w, customErr.ErrCode(), `request invalid`, nil, nil, customErr)
		return
	}
	utils.ResponseHandler(w, http.StatusOK, `success get all records`, &dataKey, result, nil)
}

var _ handlers_interfaces.TestRecordHandler = &TestRecordHandler{}

func NewTestRecordHandler(testRecordUsecase usecases_interfaces.TestRecordUseCase) *TestRecordHandler {
	return &TestRecordHandler{
		testRecordUsecase: testRecordUsecase,
	}
}

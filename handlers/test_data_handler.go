package handlers

import (
	"context"
	"github.com/google/uuid"
	"go-api-test/entities/payloads/requests"
	"go-api-test/enums"
	handlers_interfaces "go-api-test/interfaces/handlers"
	usecases_interfaces "go-api-test/interfaces/usecases"
	"go-api-test/utils"
	"net/http"
)

type TestDataHandler struct {
	usecase usecases_interfaces.TestDataUsecase
}

func (t *TestDataHandler) CreateTestAPI(w http.ResponseWriter, r *http.Request) {
	utils.ContentType(w, enums.ContentTypeJSON)
	var (
		rID     = uuid.NewString()
		ctx     = context.WithValue(r.Context(), enums.RequestID, rID)
		request = &requests.TestDataRequest{}
		dataKey = `test_result`
	)
	if err := utils.BindRequest(r, request); err != nil {
		utils.ResponseHandler(w, http.StatusBadRequest, `failed to convert bind`, nil, nil, err)
		return
	}
	result, customErr := t.usecase.CreateAPITest(ctx, request)
	if customErr != nil {
		utils.ResponseHandler(w, customErr.ErrCode(), `failed to get response`, nil, nil, customErr)
		return
	}
	utils.ResponseHandler(w, http.StatusOK, `success get response`, &dataKey, result, nil)
}

var _ handlers_interfaces.TestDataHandler = &TestDataHandler{}

func NewTestDataHandler(usecase usecases_interfaces.TestDataUsecase) *TestDataHandler {
	return &TestDataHandler{
		usecase: usecase,
	}
}

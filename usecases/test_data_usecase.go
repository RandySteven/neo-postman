package usecases

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/RandySteven/neo-postman/apperror"
	"github.com/RandySteven/neo-postman/entities/models"
	"github.com/RandySteven/neo-postman/entities/payloads/requests"
	"github.com/RandySteven/neo-postman/entities/payloads/responses"
	"github.com/RandySteven/neo-postman/enums"
	repositories_interfaces "github.com/RandySteven/neo-postman/interfaces/repositories"
	usecases_interfaces "github.com/RandySteven/neo-postman/interfaces/usecases"
	"io/ioutil"
	"net/http"
)

type testDataUsecase struct {
	testDataRepo repositories_interfaces.TestDataRepository
}

func (t *testDataUsecase) GetAllRecords(ctx context.Context) (result []*responses.TestRecordList, customErr *apperror.CustomError) {
	testDatas, err := t.testDataRepo.FindAll(ctx)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get all records`, err)
	}
	for _, testData := range testDatas {
		result = append(result, &responses.TestRecordList{
			ID:           testData.ID,
			Description:  testData.Description,
			ResultStatus: testData.ResultStatus.ToString(),
			CreatedAt:    testData.CreatedAt.Local(),
		})
	}
	return
}

func (t *testDataUsecase) CreateAPITest(ctx context.Context, request *requests.TestDataRequest) (result *responses.TestDataResponse, customErr *apperror.CustomError) {
	client := &http.Client{}
	uri := "http://localhost:8080" + request.Path
	testData := &models.TestData{
		Method:               request.Method,
		Description:          request.Description,
		RequestHeader:        request.RequestHeader,
		URI:                  uri,
		RequestBody:          request.RequestBody,
		ExpectedResponse:     request.ExpectedResponse,
		ExpectedResponseCode: request.ExpectedResponseCode,
		ActualResponse:       nil,
	}
	body, err := request.RequestBody.MarshalJSON()
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to marshal request body`, err)
	}
	req, err := http.NewRequestWithContext(ctx, request.Method, uri, ioutil.NopCloser(bytes.NewBuffer(body)))
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to hit api`, err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get response`, err)
	}

	if request.ExpectedResponse != nil {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to read response body`, err)
		}
		testData.ActualResponse = make(json.RawMessage, 0)
		err = json.Unmarshal(body, &testData.ActualResponse)
		if err != nil {
			return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to unmarshal response body`, err)
		}
	}

	for testData.ResultStatus != enums.Unexpected {
		if resp.StatusCode != request.ExpectedResponseCode {
			testData.ResultStatus = enums.Unexpected
			break
		}

		if request.ExpectedResponse != nil {
			resultStatus := enums.Expected
			for key, value := range request.ExpectedResponse {
				if testData.ActualResponse[key] != value {
					resultStatus = enums.Unexpected
				}
			}
			testData.ResultStatus = resultStatus
		}

		if testData.ResultStatus == enums.Expected {
			break
		}
	}
	testData.ActualResponseCode = resp.StatusCode

	testData, err = t.testDataRepo.Save(ctx, testData)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to insert database`, err)
	}

	result = &responses.TestDataResponse{
		ID:           testData.ID,
		ResultStatus: testData.ResultStatus.ToString(),
	}

	if testData.ResultStatus == enums.Unexpected {
		result.ExpectedResponseCode = testData.ExpectedResponseCode
		result.ActualResponseCode = testData.ActualResponseCode
		result.ExpectedResponseBody = testData.ExpectedResponse
		result.ActualResponseBody = testData.ActualResponse
	}

	return result, nil
}

var _ usecases_interfaces.TestDataUsecase = &testDataUsecase{}

func NewTestDataUsecase(testDataRepo repositories_interfaces.TestDataRepository) *testDataUsecase {
	return &testDataUsecase{
		testDataRepo: testDataRepo,
	}
}

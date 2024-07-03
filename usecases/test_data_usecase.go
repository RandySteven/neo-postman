package usecases

import (
	"context"
	"go-api-test/apperror"
	"go-api-test/entities/models"
	"go-api-test/entities/payloads/requests"
	"go-api-test/entities/payloads/responses"
	"go-api-test/enums"
	repositories_interfaces "go-api-test/interfaces/repositories"
	usecases_interfaces "go-api-test/interfaces/usecases"
	"go-api-test/utils"
	"net/http"
)

type testDataUsecase struct {
	testDataRepo repositories_interfaces.TestDataRepository
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
	body, err := utils.MapToJSONReader(request.RequestBody)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to convert request`, err)
	}
	req, err := http.NewRequestWithContext(ctx, request.Method, uri, body)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to hit api`, err)
	}
	for key, value := range testData.RequestHeader {
		req.Header.Set(key, value.(string))
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get response`, err)
	}

	if request.ExpectedResponse != nil {
		respBody, err := utils.ConvertJSON(resp)
		if err != nil {
			return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to convert response`, err)
		}
		testData.ActualResponse = respBody
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

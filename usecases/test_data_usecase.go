package usecases

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/RandySteven/neo-postman/apperror"
	"github.com/RandySteven/neo-postman/entities/models"
	"github.com/RandySteven/neo-postman/entities/payloads/requests"
	"github.com/RandySteven/neo-postman/entities/payloads/responses"
	"github.com/RandySteven/neo-postman/enums"
	repositories_interfaces "github.com/RandySteven/neo-postman/interfaces/repositories"
	usecases_interfaces "github.com/RandySteven/neo-postman/interfaces/usecases"
	"github.com/RandySteven/neo-postman/pkg/redis"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

type testDataUsecase struct {
	testDataRepo repositories_interfaces.TestDataRepository
	redis        *redis.RedisClient
}

func (t *testDataUsecase) AutoDeleteUnsavedRecord(ctx context.Context) (err error) {
	return t.testDataRepo.DeletedUnsavedTestData(ctx)
}

func (t *testDataUsecase) SaveRecord(ctx context.Context, id uint64) (result string, customErr *apperror.CustomError) {
	testData, err := t.testDataRepo.FindByID(ctx, id)
	if err != nil {
		return "", apperror.NewCustomError(apperror.ErrInternalServer, `failed to get test data`, err)
	}
	if testData.IsSaved == true {
		return "", apperror.NewCustomError(apperror.ErrBadRequest, `user try to save again`, fmt.Errorf("you already put this on record"))
	}
	testData.IsSaved = true
	testData, err = t.testDataRepo.Update(ctx, testData)
	if err != nil {
		return "", apperror.NewCustomError(apperror.ErrInternalServer, `failed to save test data`, err)
	}
	result = "success to save test data"
	return result, nil
}

func (t *testDataUsecase) GetRecord(ctx context.Context, id uint64) (result *responses.TestRecordDetail, customErr *apperror.CustomError) {
	testData, err := t.testDataRepo.FindByID(ctx, id)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get record`, err)
	}
	result = &responses.TestRecordDetail{
		ID:           testData.ID,
		Description:  testData.Description,
		Endpoint:     testData.URI[len("http://localhost:8080"):],
		Method:       testData.Method,
		ResultStatus: testData.ResultStatus.ToString(),
		ExpectedResponse: responses.TestDataExpectedResponse{
			ExpectedResponseCode: testData.ExpectedResponseCode,
			ExpectedResponse:     testData.ExpectedResponse,
		},
		ActualResponse: responses.TestDataActualResponse{
			ActualResponseCode: testData.ActualResponseCode,
			ActualResponse:     testData.ActualResponse,
		},
		CreatedAt: testData.CreatedAt.Local(),
	}
	return result, nil
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
			IsSaved:      testData.IsSaved,
			Links: struct {
				Detail  string `json:"detail"`
				Save    string `json:"save"`
				Unsaved string `json:"unsaved"`
			}{
				Detail:  fmt.Sprintf("http://localhost:8081/testdata/%d", testData.ID),
				Save:    fmt.Sprintf("http://localhost:8081/testdata/%d/saved", testData.ID),
				Unsaved: fmt.Sprintf("http://localhost:8081/testdata/%d/unsaved", testData.ID),
			},
		})
	}
	return
}

func (t *testDataUsecase) CreateAPITest(ctx context.Context, request *requests.TestDataRequest) (result *responses.TestDataResponse, customErr *apperror.CustomError) {
	start := time.Now()
	client := &http.Client{
		Transport: &http.Transport{MaxIdleConnsPerHost: 10},
	}
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
		ResultStatus:         enums.Error,
	}
	body, err := request.RequestBody.MarshalJSON()
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrBadRequest, `the request body is not valid`, err)
	}
	req, err := http.NewRequestWithContext(ctx, request.Method, uri, ioutil.NopCloser(bytes.NewBuffer(body)))
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to hit api`, err)
	}

	headerMap := make(map[string]string)
	err = json.Unmarshal(testData.RequestHeader, &headerMap)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrBadRequest, `the request body is not valid`, err)
	}
	for key, value := range headerMap {
		req.Header.Add(key, value)
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
		testData.ResultStatus = enums.Expected

		if testData.ResultStatus == enums.Expected {
			break
		}
	}
	testData.ActualResponseCode = resp.StatusCode
	var (
		resultCh    = make(chan *models.TestData, 1)
		customErrCh = make(chan *apperror.CustomError, 1)
		wg          sync.WaitGroup
	)
	wg.Add(1)
	go func() {
		defer wg.Done()
		testData, err = t.testDataRepo.Save(ctx, testData)
		if err != nil {
			customErrCh <- apperror.NewCustomError(apperror.ErrInternalServer, `failed to insert database`, err)
			close(customErrCh)
			close(resultCh)
			return
		}
		resultCh <- testData
		close(customErrCh)
		close(resultCh)
		return
	}()

	for customErr = range customErrCh {
		return nil, customErr
	}
	testData = <-resultCh
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
	end := time.Since(start)
	log.Println("latency : ", end)

	return result, nil
}

var _ usecases_interfaces.TestDataUsecase = &testDataUsecase{}

func NewTestDataUsecase(testDataRepo repositories_interfaces.TestDataRepository) *testDataUsecase {
	return &testDataUsecase{
		testDataRepo: testDataRepo,
		redis:        redis.NewRedis(),
	}
}

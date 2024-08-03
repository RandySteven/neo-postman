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
	caches_interfaces "github.com/RandySteven/neo-postman/interfaces/caches"
	repositories_interfaces "github.com/RandySteven/neo-postman/interfaces/repositories"
	usecases_interfaces "github.com/RandySteven/neo-postman/interfaces/usecases"
	"github.com/RandySteven/neo-postman/pkg/yaml"
	"github.com/RandySteven/neo-postman/utils"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type testDataUsecase struct {
	testDataRepo   repositories_interfaces.TestDataRepository
	testRecordRepo repositories_interfaces.TestRecordRepository
	testDataCache  caches_interfaces.TestDataCache
}

func (t *testDataUsecase) UnsavedRecord(ctx context.Context, id uint64) (result string, customErr *apperror.CustomError) {
	testData, err := t.testDataRepo.FindByID(ctx, id)
	if err != nil {
		return "", apperror.NewCustomError(apperror.ErrInternalServer, `failed to get test data`, err)
	}

	var (
		wg          sync.WaitGroup
		customErrCh = make(chan *apperror.CustomError)
	)

	wg.Add(2)

	go func() {
		defer wg.Done()
		if testData.IsSaved == false {
			customErrCh <- apperror.NewCustomError(apperror.ErrBadRequest, `user try to deleted again`, fmt.Errorf("you haven't put this on record"))
			return
		}
		testData.IsSaved = false
		testData, err = t.testDataRepo.Update(ctx, testData)
		if err != nil {
			customErrCh <- apperror.NewCustomError(apperror.ErrInternalServer, `failed to unsaved test data`, err)
			return
		}
	}()

	go func() {
		defer wg.Done()

	}()

	go func() {
		defer wg.Done()
		close(customErrCh)
	}()

	return
}

func (t *testDataUsecase) AutoDeleteUnsavedRecord(ctx context.Context) (err error) {
	return t.testDataRepo.DeletedUnsavedTestData(ctx)
}

func (t *testDataUsecase) SaveRecord(ctx context.Context, id uint64) (result string, customErr *apperror.CustomError) {
	testData, err := t.testDataRepo.FindByID(ctx, id)
	if err != nil {
		return "", apperror.NewCustomError(apperror.ErrInternalServer, `failed to get test data`, err)
	}

	var (
		wg          sync.WaitGroup
		customErrCh = make(chan *apperror.CustomError)
	)

	wg.Add(2)

	go func() {
		defer wg.Done()
		if testData.IsSaved == true {
			customErrCh <- apperror.NewCustomError(apperror.ErrBadRequest, `user try to save again`, fmt.Errorf("you already put this on record"))
			return
		}
		testData.IsSaved = true
		testData, err = t.testDataRepo.Update(ctx, testData)
		if err != nil {
			customErrCh <- apperror.NewCustomError(apperror.ErrInternalServer, `failed to save test data`, err)
			return
		}
	}()

	go func() {
		defer wg.Done()

		_, err = t.testRecordRepo.Save(ctx, &models.TestRecord{TestDataID: testData.ID})
		if err != nil {
			customErrCh <- apperror.NewCustomError(apperror.ErrInternalServer, `failed to save test record`, err)
			return
		}
	}()

	go func() {
		wg.Wait()
		close(customErrCh)
	}()

	select {
	case customErr = <-customErrCh:
		return "", customErr
	default:
		result = "success to save test data"
		return result, nil
	}
}

func (t *testDataUsecase) GetRecord(ctx context.Context, id uint64) (result *responses.TestDataDetail, customErr *apperror.CustomError) {
	testData, err := t.testDataCache.Get(ctx, strconv.Itoa(int(id)))
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get test data cache`, err)
	}
	testData, err = t.testDataRepo.FindByID(ctx, id)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get record`, err)
	}
	err = t.testDataCache.Set(ctx, strconv.Itoa(int(id)), testData)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to save test data cache`, err)
	}
	result = &responses.TestDataDetail{
		ID:           testData.ID,
		Description:  testData.Description,
		Endpoint:     testData.URI,
		IsSaved:      testData.IsSaved,
		Method:       testData.Method,
		ResultStatus: testData.ResultStatus.ToString(),
		ExpectedResponse: responses.TestDataExpectedResponse{
			ExpectedResponseCode: testData.ExpectedResponseCode,
			ExpectedResponse:     testData.ExpectedResponse,
		},
		RequestHeader: testData.RequestHeader,
		RequestBody:   testData.RequestBody,
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
		detailUrl := utils.DetailURL(enums.TestDataPrefix.ToString(), testData.ID)
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
				Detail:  detailUrl,
				Save:    detailUrl + "/saved",
				Unsaved: detailUrl + "/unsaved",
			},
		})
	}
	return
}

func (t *testDataUsecase) CreateAPITest(ctx context.Context, request *requests.TestDataRequest) (result *responses.TestDataResponse, customErr *apperror.CustomError) {
	start := time.Now()
	baseUrl, err := yaml.ReadBaseURLYAML()
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed get base url`, err)
	}
	exists := baseUrl.CheckURLExists(request.URLKey)
	if !exists {
		return nil, apperror.NewCustomError(apperror.ErrBadRequest, `url doesn exists`, fmt.Errorf("url doesn exists"))
	}
	client := &http.Client{
		Transport: &http.Transport{MaxIdleConnsPerHost: 10},
	}
	uri := baseUrl.UrlList[request.URLKey] + request.Path
	testData := &models.TestData{
		Method:               request.Method,
		Description:          request.Description,
		RequestHeader:        request.RequestHeader,
		Host:                 baseUrl.UrlList[request.URLKey],
		URI:                  request.Path,
		RequestBody:          request.RequestBody,
		ExpectedResponse:     request.ExpectedResponse,
		ExpectedResponseCode: request.ExpectedResponseCode,
		ActualResponse:       nil,
		ResultStatus:         enums.Error,
	}
	var (
		req                  = &http.Request{}
		reader io.ReadCloser = nil
	)
	if request.RequestBody != nil {
		body, err := request.RequestBody.MarshalJSON()
		if err != nil {
			return nil, apperror.NewCustomError(apperror.ErrBadRequest, `the request body is not valid`, err)
		}
		reader = ioutil.NopCloser(bytes.NewBuffer(body))
	}
	log.Println("masuk header")
	req, err = http.NewRequestWithContext(ctx, request.Method, uri, reader)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to hit api`, err)
	}

	headerMap := make(map[string]string)
	err = json.Unmarshal(testData.RequestHeader, &headerMap)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrBadRequest, `the request body is not valid`, err)
	}
	log.Println("header map ", headerMap)
	for key, value := range headerMap {
		req.Header.Add(key, value)
	}

	responseTimeStart := time.Now()
	resp, err := client.Do(req)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get response`, err)
	}
	responseTime := time.Since(responseTimeStart)
	testData.ResponseTime = responseTime

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

		// Compare actual response with expected response
		if !compareJSON(testData.ActualResponse, request.ExpectedResponse) {
			testData.ResultStatus = enums.Unexpected
		} else {
			testData.ResultStatus = enums.Expected
		}
	}

	testData.ActualResponseCode = resp.StatusCode

	// Save test data to database
	savedTestData, err := t.testDataRepo.Save(ctx, testData)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to insert database`, err)
	}

	detailUrl := utils.DetailURL(enums.TestDataPrefix.ToString(), savedTestData.ID)
	result = &responses.TestDataResponse{
		ID:           savedTestData.ID,
		ResultStatus: savedTestData.ResultStatus.ToString(),
		Links: struct {
			Detail string `json:"detail"`
			Saved  string `json:"saved"`
		}{
			Detail: detailUrl,
			Saved:  detailUrl + "/saved",
		},
		ResponseTime:         responseTime,
		ExpectedResponseCode: savedTestData.ExpectedResponseCode,
		ActualResponseCode:   savedTestData.ExpectedResponseCode,
		ExpectedResponseBody: savedTestData.ExpectedResponse,
		ActualResponseBody:   savedTestData.ActualResponse,
	}

	end := time.Since(start)
	log.Println("latency : ", end)

	return result, nil
}

func compareJSON(actual, expected json.RawMessage) bool {
	var actualData, expectedData map[string]interface{}
	if err := json.Unmarshal(actual, &actualData); err != nil {
		return false
	}
	if err := json.Unmarshal(expected, &expectedData); err != nil {
		return false
	}

	if actualData["data"] != nil && expectedData["data"] != nil {
		actualCaptureMethod := actualData["data"].(map[string]interface{})["captureMethod"]
		expectedCaptureMethod := expectedData["data"].(map[string]interface{})["captureMethod"]

		if actualCaptureMethod != expectedCaptureMethod {
			return false
		}
	}

	return true
}

func jsonToMap(something json.RawMessage) (result map[string]interface{}) {
	err := json.Unmarshal(something, &result)
	if err != nil {
		return
	}
	return result
}

var _ usecases_interfaces.TestDataUsecase = &testDataUsecase{}

func NewTestDataUsecase(
	testDataRepo repositories_interfaces.TestDataRepository,
	testRecordRepo repositories_interfaces.TestRecordRepository,
	testDataCache caches_interfaces.TestDataCache,
) *testDataUsecase {
	return &testDataUsecase{
		testDataRepo:   testDataRepo,
		testRecordRepo: testRecordRepo,
		testDataCache:  testDataCache,
	}
}

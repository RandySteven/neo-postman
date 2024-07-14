package usecases

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/RandySteven/neo-postman/apperror"
	"github.com/RandySteven/neo-postman/entities/models"
	"github.com/RandySteven/neo-postman/entities/payloads/requests"
	"github.com/RandySteven/neo-postman/entities/payloads/responses"
	repositories_interfaces "github.com/RandySteven/neo-postman/interfaces/repositories"
	usecases_interfaces "github.com/RandySteven/neo-postman/interfaces/usecases"
	"log"
	"sync"
	"time"
)

type testRecordUsecase struct {
	testDataRepo   repositories_interfaces.TestDataRepository
	testRecordRepo repositories_interfaces.TestRecordRepository
}

func (t *testRecordUsecase) AutoSaveTestRecord(ctx context.Context) (err error) {
	return
}

func (t *testRecordUsecase) CreateTestRecord(ctx context.Context, request *requests.TestRecordRequest) (customErr *apperror.CustomError) {
	testRecord := &models.TestRecord{}
	if len(request.TestDataIDs) != 0 {

	} else {
		testRecord.TestDataID = request.TestDataID
		_, err := t.testRecordRepo.Save(ctx, testRecord)
		if err != nil {
			return apperror.NewCustomError(apperror.ErrInternalServer, `failed to create test record`, err)
		}
	}
	return nil
}

func (t *testRecordUsecase) GetAllTestRecords(ctx context.Context) (result []*responses.TestRecordListResponse, customErr *apperror.CustomError) {
	startTime := time.Now()
	testRecords, err := t.testRecordRepo.FindAll(ctx)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get test records`, err)
	}
	customErrCh := make(chan *apperror.CustomError)
	resultCh := make(chan []*responses.TestRecordListResponse)
	wg := sync.WaitGroup{}

	wg.Add(len(testRecords))

	go func() {
		defer wg.Done()
		for _, testRecord := range testRecords {
			testData, err := t.testDataRepo.FindByID(ctx, testRecord.TestDataID)
			if err != nil {
				customErr = apperror.NewCustomError(apperror.ErrInternalServer, `failed to get test data`, err)
				customErrCh <- customErr
				return
			}
			result = append(result, &responses.TestRecordListResponse{
				ID:          testRecord.ID,
				Endpoint:    testData.URI,
				Description: testData.Description,
				Links: struct {
					Detail string `json:"detail"`
				}{
					Detail: fmt.Sprintf("http://localhost:8081/testrecord/%d", testRecord.ID),
				},
				CreatedAt: testRecord.CreatedAt,
				UpdatedAt: testRecord.UpdatedAt,
			})
		}
		resultCh <- result
		return
	}()

	go func() {
		wg.Wait()
		close(customErrCh)
		close(resultCh)
	}()

	log.Println("latency : ", time.Since(startTime))

	select {
	case customErr = <-customErrCh:
		return nil, customErr
	case result = <-resultCh:
		return result, nil
	}
}

func (t *testRecordUsecase) GetTestRecordDetail(ctx context.Context, id uint64) (result *responses.TestRecordDetailResponse, customErr *apperror.CustomError) {
	testRecord, err := t.testRecordRepo.FindByID(ctx, id)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get test record`, err)
	}
	testData, err := t.testDataRepo.FindByID(ctx, testRecord.TestDataID)
	if err != nil {
		return nil, apperror.NewCustomError(apperror.ErrInternalServer, `failed to get test data`, err)
	}
	result = &responses.TestRecordDetailResponse{
		ID: testRecord.ID,
		TestData: struct {
			ID            uint64          `json:"id"`
			RequestHeader json.RawMessage `json:"request_header"`
			RequestBody   json.RawMessage `json:"request_body"`
			ResultStatus  string          `json:"result_status"`
			ResponseCode  int             `json:"response_code"`
			Links         struct {
				Detail string `json:"detail"`
			} `json:"links"`
		}{
			ID:            testData.ID,
			RequestHeader: testData.RequestHeader,
			RequestBody:   testData.RequestBody,
			ResultStatus:  testData.ResultStatus.ToString(),
			ResponseCode:  testData.ActualResponseCode,
			Links: struct {
				Detail string `json:"detail"`
			}{Detail: fmt.Sprintf("http://localhost:8081/testdata/%d", testData.ID)},
		},
		CreatedAt: testData.CreatedAt,
		UpdatedAt: testData.UpdatedAt,
	}
	return result, nil
}

var _ usecases_interfaces.TestRecordUseCase = &testRecordUsecase{}

func NewTestRecordUsecase(
	testDataRepo repositories_interfaces.TestDataRepository,
	testRecordRepo repositories_interfaces.TestRecordRepository,
) *testRecordUsecase {
	return &testRecordUsecase{
		testDataRepo:   testDataRepo,
		testRecordRepo: testRecordRepo,
	}
}

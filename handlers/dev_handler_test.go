package handlers_test

import (
	"github.com/RandySteven/neo-postman/handlers"
	"github.com/RandySteven/neo-postman/pkg/yaml"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type DevHandlerTestSuite struct {
	suite.Suite
}

func (suite *DevHandlerTestSuite) SetupTest() {

}

func (suite *DevHandlerTestSuite) TestHello() {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		suite.T().Error(err)
	}
	rr := httptest.NewRecorder()
	var devHandler = &handlers.DevHandler{}

	handler := http.HandlerFunc(devHandler.Hello)
	handler.ServeHTTP(rr, req)
	suite.Equal(http.StatusOK, rr.Code)
}

func (suite *DevHandlerTestSuite) TestGetListUrl() {
	req, err := http.NewRequest("GET", "/dev/listurl", nil)
	if err != nil {
		suite.T().Error(err)
	}
	_, err = yaml.ReadBaseURLYAML()
	if err != nil {
		suite.T().Error(err)
	}
	rr := httptest.NewRecorder()
	var devHandler = &handlers.DevHandler{}

	handler := http.HandlerFunc(devHandler.GetListUrl)
	handler.ServeHTTP(rr, req)
	suite.Equal(http.StatusOK, rr.Code)
}

func TestDevHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(DevHandlerTestSuite))
}

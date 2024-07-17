package handlers

import (
	handlers_interfaces "github.com/RandySteven/neo-postman/interfaces/handlers"
	"github.com/RandySteven/neo-postman/pkg/yaml"
	"github.com/RandySteven/neo-postman/utils"
	"net/http"
)

type DevHandler struct{}

func (d *DevHandler) GetListUrl(w http.ResponseWriter, r *http.Request) {
	readBaseUrl, err := yaml.ReadBaseURLYAML()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	dataKey := `baseUrl`
	utils.ResponseHandler(w, http.StatusOK, `success get base url`, &dataKey, readBaseUrl.UrlList, nil)
}

func (d *DevHandler) DummyTester(w http.ResponseWriter, r *http.Request) {
}

func (d *DevHandler) Hello(w http.ResponseWriter, r *http.Request) {
}

var _ handlers_interfaces.DevHandler = &DevHandler{}

func NewDevHandler() *DevHandler {
	return &DevHandler{}
}

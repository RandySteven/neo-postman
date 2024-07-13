package utils

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/RandySteven/neo-postman/apperror"
	"github.com/RandySteven/neo-postman/entities/payloads/responses"
	"github.com/RandySteven/neo-postman/enums"
	"github.com/iancoleman/strcase"
	"io"
	"log"
	"net/http"
	"reflect"
)

func ContentType(w http.ResponseWriter, contentType string) {
	w.Header().Set("Content-Type", contentType)
}

func BindJSON(req *http.Request, request interface{}) error {
	return json.NewDecoder(req.Body).Decode(&request)
}

func BindXML(req *http.Request, request interface{}) error {
	return xml.NewDecoder(req.Body).Decode(&request)
}

func BindForm(req *http.Request, request interface{}) error {
	err := req.ParseForm()
	if err != nil {
		return err
	}

	requestValue := reflect.ValueOf(request)
	requestElem := requestValue.Elem()

	formValues := make(map[string]reflect.Value)

	for key, values := range req.Form {
		snakeKey := strcase.ToSnake(key)
		field := requestElem.FieldByName(snakeKey)
		if !field.IsValid() {
			log.Printf("Field not found in request: %s", snakeKey)
			continue
		}
		if !field.CanSet() {
			return fmt.Errorf("Field %s is not settable", snakeKey)
		}
		formValue := reflect.ValueOf(values[0])
		if !formValue.Type().ConvertibleTo(field.Type()) {
			return fmt.Errorf("Type conversion error for field %s", snakeKey)
		}
		convertedValue := formValue.Convert(field.Type())
		field.Set(convertedValue)
		formValues[snakeKey] = convertedValue
	}

	return nil
}

func BindMultipartForm(req *http.Request, request interface{}) error {
	err := req.ParseMultipartForm(32 << 20) // Adjust max memory size as needed
	if err != nil {
		return err
	}
	err = BindForm(req, request)
	if err != nil {
		return err
	}
	multipartFiles := req.MultipartForm.File

	for fieldname, files := range multipartFiles {
		field := reflect.ValueOf(request).Elem().FieldByName(fieldname)
		if !field.IsValid() {
			log.Printf("Field not found in request for file upload: %s", fieldname)
			continue
		}

		if field.Kind() != reflect.Slice || !reflect.TypeOf(field.Interface()).Elem().Implements(reflect.TypeOf((*io.Reader)(nil)).Elem()) {
			return fmt.Errorf("Field %s is not a slice of io.Reader", fieldname)
		}

		fieldSlice := reflect.MakeSlice(field.Type(), 0, len(files))

		for _, fileHeader := range files {
			uploadedFile, err := fileHeader.Open()
			if err != nil {
				return fmt.Errorf("Error opening uploaded file %s: %w", fileHeader.Filename, err)
			}
			defer uploadedFile.Close() // Close the file after processing

			fieldSlice = reflect.Append(fieldSlice, reflect.ValueOf(uploadedFile))
		}

		field.Set(fieldSlice)
	}

	return nil
}

func ResponseHandler(w http.ResponseWriter, responseCode int, message string, dataKey *string, responseData any, err error) {
	ContentType(w, "application/json")
	w.WriteHeader(responseCode)
	responseMap := make(map[string]any)
	if dataKey != nil && responseData != nil {
		responseMap[*dataKey] = responseData
	}
	response := responses.NewResponse(message, responseMap, err)
	log.Print(response)
	err = json.NewEncoder(w).Encode(&response)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func BindRequest(req *http.Request, request interface{}) error {
	var bindRequestType = map[string]error{
		enums.ContentTypeJSON:     BindJSON(req, request),
		enums.ContentTypeForm:     BindForm(req, request),
		enums.ContentTypeFormData: BindMultipartForm(req, request),
	}

	contentType := req.Header.Get("Content-Type")

	return bindRequestType[contentType]
}

func ErrorHandler(w http.ResponseWriter, customErr *apperror.CustomError) {
	w.WriteHeader(customErr.ErrCode())
	response := responses.NewResponse("", nil, customErr)
	json.NewEncoder(w).Encode(&response)
}

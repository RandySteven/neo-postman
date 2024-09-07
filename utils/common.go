package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func ConvertJSON(reader *http.Response) (map[string]interface{}, error) {
	defer reader.Body.Close() // Ensure closing the reader

	// Read all bytes from the ReadCloser
	log.Println("reader : ", reader)
	body, err := ioutil.ReadAll(reader.Body)
	if err != nil {
		return nil, err
	}
	log.Println("body : ", body)

	// Create a map to mongos the decoded data
	result := make(map[string]interface{})

	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Println("error unmarshalling body ", err)
		return nil, err
	}

	return result, nil
}

func DetailURL(prefix string, id uint64) string {
	return fmt.Sprintf("%s/%s/%d", os.Getenv("APP_HOST"), prefix, id)
}

func MapToJSONReader(data map[string]interface{}) (io.Reader, error) {
	// Encode the map to JSON bytes
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// Create a reader from the JSON byte slice
	return bytes.NewReader(jsonData), nil
}

func JsonString(request map[string]interface{}) (string, error) {
	bytes, err := json.Marshal(request)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

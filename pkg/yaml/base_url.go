package yaml

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

type BaseUrl struct {
	UrlList map[string]string `yaml:"urlList"`
}

func ReadBaseURLYAML() (*BaseUrl, error) {

	fileName, err := filepath.Abs("./files/yml/baseUrl.local.yml")
	if err != nil {
		return nil, err
	}

	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	baseUrl := &BaseUrl{}

	err = yaml.Unmarshal(yamlFile, &baseUrl)
	if err != nil {
		return nil, err
	}

	return baseUrl, nil
}

func (b *BaseUrl) CheckURLExists(key string) bool {
	if b.UrlList[key] != "" {
		return true
	}
	return false
}

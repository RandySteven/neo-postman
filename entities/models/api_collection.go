package models

import "time"

type (
	ApiCollection struct {
		ID          uint64
		Title       string
		Description string
		Collection  string
		CreatedAt   time.Time
		UpdatedAt   time.Time
		DeletedAt   *time.Time
	}

	ApiCollectionDetail struct {
		Info struct {
			PostManId      string `json:"_postman_id" bson:"_postman_id"`
			Name           string `json:"name" bson:"name"`
			Schema         string `json:"schema" bson:"schema"`
			ExporterId     string `json:"_exporter_id" bson:"_exporter_id"`
			CollectionLink string `json:"_collection_link" bson:"_collection_link"`
		} `json:"info" bson:"info"`
		Item []struct {
			Name    string `json:"name" bson:"name"`
			Request struct {
				Method string `json:"method" bson:"method"`
				Header []struct {
					Key   string `json:"key" bson:"key"`
					Value string `json:"value" bson:"value"`
					Type  string `json:"type" bson:"type"`
				} `json:"header" bson:"header"`
				Body struct {
					Mode    string `json:"mode" bson:"mode"`
					Raw     string `json:"raw" bson:"raw"`
					Options struct {
						Raw struct {
							Language string `json:"language" bson:"language"`
						} `json:"raw" bson:"raw"`
					} `json:"options" bson:"options"`
				} `json:"body" bson:"body"`
				Url struct {
					Raw  string   `json:"raw" bson:"raw"`
					Host []string `json:"host" bson:"host"`
					Path []string `json:"path" bson:"path"`
				} `json:"url" bson:"url"`
			}
			Responses []struct{} `json:"responses" bson:"responses"`
		}
	}
)

package documentaries

import "github.com/RandySteven/neo-postman/entities/models"

type (
	MultiMatch struct {
		Query  string   `json:"query"`
		Fields []string `json:"fields"`
		Type   string   `json:"type"`
	}

	Query struct {
		MultiMatch MultiMatch `json:"multi_match"`
	}

	SearchRequest struct {
		Query Query `json:"query"`
	}

	SearchResponse struct {
		Hits struct {
			Hits []struct {
				Source models.TestData `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}
)

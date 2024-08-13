package responses

import "time"

type (
	APICollectionCreateResponse struct {
		ID         uint64    `json:"id"`
		Title      string    `json:"title"`
		Collection string    `json:"collection"`
		CreatedAt  time.Time `json:"created_at"`
	}

	APICollectionListResponse struct {
		ID        uint64    `json:"id"`
		Title     string    `json:"title"`
		CreatedAt time.Time `json:"created_at"`
	}

	APICollectionDetailResponse struct {
		ID          uint64     `json:"id"`
		Title       string     `json:"title"`
		Description string     `json:"description"`
		Collection  string     `json:"collection"`
		CreatedAt   time.Time  `json:"created_at"`
		UpdatedAt   time.Time  `json:"updated_at"`
		DeletedAt   *time.Time `json:"deleted_at"`
	}
)

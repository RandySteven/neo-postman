package requests

type APICollectionRequest struct {
	Title          string `json:"title"`
	CollectionFile string `form:"collection_file"`
}

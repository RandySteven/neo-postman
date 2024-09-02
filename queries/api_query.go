package queries

const (
	InsertApiQuery GoQuery = `
		INSERT INTO apis (title, description, content_file) 
		VALUES ($1, $2, $3)
	`

	SelectApisQuery GoQuery = `
		SELECT id, title, description, content_file, created_at, updated_at, deleted_at 
		FROM apis
		WHERE deleted_at IS NULL
	`

	SelectApiByID GoQuery = `
		SELECT id, title, description, content_file, created_at, updated_at, deleted_at
		FROM apis
		WHERE id = $1
	`
)

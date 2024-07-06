package repositories

import (
	"context"
	"database/sql"
	"github.com/RandySteven/neo-postman/entities/models"
	repositories_interfaces "github.com/RandySteven/neo-postman/interfaces/repositories"
	"github.com/RandySteven/neo-postman/queries"
	"github.com/RandySteven/neo-postman/utils"
)

type jiraIssueRepository struct {
	db *sql.DB
}

func (j *jiraIssueRepository) Save(ctx context.Context, request *models.JiraIssue) (result *models.JiraIssue, err error) {
	id, err := utils.Save[models.JiraIssue](ctx, j.db, queries.InsertJiraIssue,
		&request.Request, &request.Response, &request.Link, &request.CreatedBy, &request.UpdatedBy)
	if err != nil {
		return nil, err
	}
	result = request
	result.ID = *id
	return result, nil
}

func (j *jiraIssueRepository) FindAll(ctx context.Context) (result []*models.JiraIssue, err error) {
	return utils.FindAll[models.JiraIssue](ctx, j.db, queries.SelectAllJiraIssues)
}

func (j *jiraIssueRepository) FindByID(ctx context.Context, id uint64) (result *models.JiraIssue, err error) {
	return
}

func (j *jiraIssueRepository) Update(ctx context.Context, request *models.JiraIssue) (result *models.JiraIssue, err error) {
	return
}

func (j *jiraIssueRepository) Delete(ctx context.Context, id uint64) (err error) {
	return
}

var _ repositories_interfaces.JiraIssueRepository = &jiraIssueRepository{}

func NewJiraRepository(db *sql.DB) *jiraIssueRepository {
	return &jiraIssueRepository{
		db: db,
	}
}

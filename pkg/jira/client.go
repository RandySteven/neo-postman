package jira

import (
	"context"
	"github.com/andygrunwald/go-jira"
)

type jiraClient struct {
	client *jira.Client
}

func (j *jiraClient) CreateIssue(ctx context.Context, request *jira.Issue) (response *jira.Response, err error) {
	_, response, err = j.client.Issue.Create(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func NewJiraClient() (*jiraClient, error) {
	client, _ := jira.NewClient(nil, "")
	return &jiraClient{
		client: client,
	}, nil
}

var _ JiraAction = &jiraClient{}

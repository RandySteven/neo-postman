package jira

import (
	"context"
	"github.com/andygrunwald/go-jira"
	"log"
	"os"
)

type jiraClient struct {
	client *jira.Client
}

func (j *jiraClient) CreateIssue(ctx context.Context, request *jira.Issue) (response *jira.Response, err error) {
	_, response, err = j.client.Issue.Create(request)
	if err != nil {
		log.Println("Error creating issue:", err.Error())
		return nil, err
	}
	return response, nil
}

func NewJiraClient() (*jiraClient, error) {
	tp := jira.BasicAuthTransport{
		Password: os.Getenv("JIRA_TOKEN"),
	}
	client, err := jira.NewClient(tp.Client(), "https://bulletinboard.atlassian.net/")
	if err != nil {
		return nil, err
	}
	return &jiraClient{
		client: client,
	}, nil
}

var _ JiraAction = &jiraClient{}

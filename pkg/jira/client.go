package jira

import (
	"context"
	"errors"
	"fmt"
	"github.com/andygrunwald/go-jira"
	"log"
	"os"
)

type jiraClient struct {
	client *jira.Client
}

func (j *jiraClient) GetClient() *jira.Client {
	return j.client
}

func (j *jiraClient) CreateIssue(ctx context.Context, request *jira.Issue) (response *jira.Response, err error) {
	requestStr := fmt.Sprintf("%v", request)
	log.Println("jira issue : ", requestStr)
	_, response, err = j.client.Issue.Create(request)
	if err != nil {
		log.Println("Error creating issue:", err.Error())
		return nil, err
	}
	return response, nil
}

func NewJiraClient() (*jiraClient, error) {
	jiraToken := os.Getenv("JIRA_TOKEN")
	if jiraToken == "" {
		return nil, errors.New("JIRA_TOKEN environment variable not set")
	}

	tp := jira.BasicAuthTransport{
		Password: jiraToken,
	}
	client, err := jira.NewClient(tp.Client(), "https://bulletinboard.atlassian.net")
	if err != nil {
		return nil, err
	}
	return &jiraClient{
		client: client,
	}, nil
}

var _ JiraAction = &jiraClient{}

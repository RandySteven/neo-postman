package jira

import "github.com/andygrunwald/go-jira"

type JiraClient struct {
	client *jira.Client
}

func NewJiraClient() (*JiraClient, error) {
	jiraClient, _ := jira.NewClient(nil, "")
	return &JiraClient{
		client: jiraClient,
	}, nil
}

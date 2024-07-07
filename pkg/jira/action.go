package jira

import (
	"context"
	"github.com/andygrunwald/go-jira"
)

type JiraAction interface {
	CreateIssue(ctx context.Context, request *jira.Issue) (response *jira.Response, err error)
}

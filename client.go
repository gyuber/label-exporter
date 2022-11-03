package exporter

import (
	"context"
	"errors"
	"os"

	"github.com/google/go-github/v48/github"
	"golang.org/x/oauth2"
)

type githubClient struct {
	client *github.Client
}

func NewClient() (*githubClient, error) {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		return nil, errors.New("missing GITHUB_TOKEN")
	}
	
	baseURL := os.Getenv("BASE_URL")
	uploadURL := os.Getenv("UPLOAD_URL")
	if baseURL != "" && uploadURL != "" {
		cli := newEnterpriseClient(baseURL, uploadURL, token)
		return &githubClient{
			client: cli,
		}, nil
	} else {
		cli := newClient(token)
		return &githubClient{
			client: cli,
		}, nil 
	}
}

func newClient(token string) *github.Client {
	ts := oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: token,
	})
	tc := oauth2.NewClient(context.Background(), ts)
	return github.NewClient(tc)
}

func newEnterpriseClient(baseURL, uploadURL string, token string) *github.Client {
    ts := oauth2.StaticTokenSource(&oauth2.Token{
        AccessToken: token,
    })
    tc := oauth2.NewClient(context.Background(), ts)
    client, _ := github.NewEnterpriseClient(baseURL, uploadURL, tc)
    return client
}

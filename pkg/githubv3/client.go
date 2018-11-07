package githubv3

import (
	"net/http"

	"github.com/google/go-github/github"
)

type Client struct {
	client *github.Client
}

func NewClient(httpClient *http.Client) *Client {
	return &Client{
		client: github.NewClient(httpClient),
	}
}

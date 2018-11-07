package githubv4

import (
	"net/http"

	"github.com/shurcooL/githubv4"
)

type Client struct {
	client *githubv4.Client
}

func NewClient(httpClient *http.Client) *Client {
	return &Client{
		client: githubv4.NewClient(httpClient),
	}
}

package githubv3

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
)

func (c Client) Search(ctx context.Context, query, org string, size int) ([]string, error) {
	var names []string

	if org != "" {
		query = fmt.Sprintf("org:%s %s", org, query)
	}

	options := &github.SearchOptions{
		Sort: "updated",
		ListOptions: github.ListOptions{
			PerPage: size,
		},
	}
	results, _, err := c.client.Search.Repositories(ctx, query, options)
	if err != nil {
		return nil, err
	}

	for _, repository := range results.Repositories {
		names = append(names, *repository.Name)
	}

	return names, nil
}

package githubv4

import (
	"context"
	"fmt"

	"github.com/shurcooL/githubv4"
)

func (c Client) Search(ctx context.Context, query, org string, size int) ([]string, error) {
	var names []string

	var graphQuery struct {
		Search struct {
			Nodes []struct {
				Repository struct {
					Name string
				} `graphql:"... on Repository"`
			}
		} `graphql:"search(type: REPOSITORY, query: $query, first: $size)"`
	}

	if org != "" {
		query = fmt.Sprintf("org:%s %s", org, query)
	}

	variables := map[string]interface{}{
		"query": githubv4.String(query),
		"size":  githubv4.Int(size),
	}
	err := c.client.Query(ctx, &graphQuery, variables)
	if err != nil {
		return nil, err
	}

	for _, node := range graphQuery.Search.Nodes {
		names = append(names, node.Repository.Name)
	}

	return names, nil
}

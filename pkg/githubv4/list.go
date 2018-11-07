package githubv4

import (
	"context"

	"github.com/shurcooL/githubv4"
)

func (c Client) List(ctx context.Context, org string, size int) ([]string, error) {
	var names []string

	var graphQuery struct {
		Organization struct {
			Repositories struct {
				PageInfo struct {
					EndCursor   string
					HasNextPage bool
				}
				Nodes []struct {
					Name string
				}
			} `graphql:"repositories(first: $size, after: $cursor)"`
		} `graphql:"organization(login: $organization)"`
	}

	variables := map[string]interface{}{
		"organization": githubv4.String(org),
		"size":         githubv4.Int(size),
		"cursor":       (*githubv4.String)(nil),
	}

	for {
		err := c.client.Query(ctx, &graphQuery, variables)
		if err != nil {
			return nil, err
		}

		for _, node := range graphQuery.Organization.Repositories.Nodes {
			names = append(names, node.Name)
		}
		if !graphQuery.Organization.Repositories.PageInfo.HasNextPage {
			break
		}
		variables["cursor"] = githubv4.String(graphQuery.Organization.Repositories.PageInfo.EndCursor)
	}

	return names, nil
}

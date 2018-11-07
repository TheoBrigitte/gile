package githubv3

import (
	"context"

	"github.com/google/go-github/github"
)

func listByOrg(client *github.Client, org string, ctx context.Context, size, next int) ([]*github.Repository, *github.Response, error) {
	opt := &github.RepositoryListByOrgOptions{
		ListOptions: github.ListOptions{
			PerPage: 100,
			Page:    next,
		},
	}
	return client.Repositories.ListByOrg(ctx, org, opt)
}

func (c Client) List(ctx context.Context, org string, size int) ([]string, error) {
	var names []string

	// list all repositories for the authenticated user
	/*repos, _, err := client.Repositories.List(ctx, "", nil)
	  if err != nil {
	          die(err)
	  }*/

	// first batch
	var next int
	var allRepos []*github.Repository
	repos, resp, err := listByOrg(c.client, org, ctx, size, next)
	if err != nil {
		return nil, err
	}
	allRepos = append(allRepos, repos...)

	// the rest in parallel
	results := make(chan []*github.Repository)
	errors := make(chan error)

	for i := resp.NextPage; i <= resp.LastPage; i++ {
		//repos, resp, err := client.Repositories.ListByOrg(ctx, *org, opt)
		go func(i int) {
			repos, _, err := listByOrg(c.client, org, ctx, size, i)
			if err != nil {
				errors <- err
			} else {
				results <- repos
			}
		}(i)
	}

	for i := resp.NextPage; i <= resp.LastPage; i++ {
		select {
		case err := <-errors:
			return nil, err
		default:
		}

		res := <-results
		allRepos = append(allRepos, res...)
	}

	for _, r := range allRepos {
		names = append(names, *r.Name)
	}

	return names, nil
}

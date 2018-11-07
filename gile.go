package main // import "github.com/TheoBrigitte/gile"

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func die(v ...interface{}) {
	fmt.Print(v)
	os.Exit(1)
}

func printRepos(repos []*github.Repository) {
	for _, r := range repos {
		fmt.Println(*r.Name)
	}
}

func listByOrg(client *github.Client, org string, ctx context.Context, next int) ([]*github.Repository, *github.Response, error) {
	opt := &github.RepositoryListByOrgOptions{
		ListOptions: github.ListOptions{
			PerPage: 100,
			Page:    next,
		},
	}
	return client.Repositories.ListByOrg(ctx, org, opt)
}

func main() {
	org := flag.String("org", "", "github organization")
	flag.Parse()

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// list all repositories for the authenticated user
	/*repos, _, err := client.Repositories.List(ctx, "", nil)
	if err != nil {
		die(err)
	}*/

	// first batch
	var next int
	var allRepos []*github.Repository
	repos, resp, err := listByOrg(client, *org, ctx, next)
	if err != nil {
		die(err)
	}
	allRepos = append(allRepos, repos...)

	// the rest in parallel
	results := make(chan []*github.Repository)
	errors := make(chan error)

	for i := resp.NextPage; i <= resp.LastPage; i++ {
		//repos, resp, err := client.Repositories.ListByOrg(ctx, *org, opt)
		go func(i int) {
			repos, _, err := listByOrg(client, *org, ctx, i)
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
			die(err)
		default:
		}

		res := <-results
		allRepos = append(allRepos, res...)
	}

	printRepos(allRepos)
	//fmt.Printf("%d repositories\n", len(allRepos))
}

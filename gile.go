package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func main() {
	org := flag.String("org", "", "github organization name")
	query := flag.String("query", "", "search query")
	size := flag.Int("size", 10, "number of items to return")
	flag.Parse()

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	if *org != "" {
		*query = fmt.Sprintf("org:%s %s", *org, *query)
	}

	opt := &github.SearchOptions{
		ListOptions: github.ListOptions{
			PerPage: *size,
		},
	}
	results, _, err := client.Search.Repositories(ctx, *query, opt)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, r := range results.Repositories {
		fmt.Println(*r.Name)
	}
}

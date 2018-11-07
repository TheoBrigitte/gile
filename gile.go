package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/TheoBrigitte/gile/pkg/github"
	"github.com/TheoBrigitte/gile/pkg/githubv3"
	"github.com/TheoBrigitte/gile/pkg/githubv4"
	"github.com/TheoBrigitte/gile/pkg/oauth2"
)

func main() {
	var (
		org        = flag.String("org", "", "github organization name")
		query      = flag.String("query", "", "search query")
		size       = flag.Int("size", 10, "number of items to return")
		apiVersion = flag.String("api", "v3", "github API version")
	)
	flag.Parse()

	ctx := context.Background()
	tc := oauth2.NewClientWithToken(ctx, os.Getenv("GITHUB_TOKEN"))

	var client github.Searcher
	switch *apiVersion {
	case "v3":
		client = githubv3.NewClient(tc)
	case "v4":
		client = githubv4.NewClient(tc)
	default:
		fmt.Printf("unsupported api version %#q", *apiVersion)
		os.Exit(1)
	}

	results, err := client.Search(ctx, *query, *org, *size)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, name := range results {
		fmt.Println(name)
	}
}

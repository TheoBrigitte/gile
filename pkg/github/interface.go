package github

import (
	"context"
)

type Lister interface {
	List(context.Context, string, int) ([]string, error)
}

type Searcher interface {
	Search(context.Context, string, string, int) ([]string, error)
}

type ListSearcher interface {
	Lister
	Searcher
}

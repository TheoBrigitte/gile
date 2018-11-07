package github

import (
	"context"
)

type Searcher interface {
	Search(context.Context, string, string, int) ([]string, error)
}

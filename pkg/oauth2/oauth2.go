package oauth2

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"
)

func NewClientWithToken(ctx context.Context, token string) *http.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	return oauth2.NewClient(ctx, ts)
}

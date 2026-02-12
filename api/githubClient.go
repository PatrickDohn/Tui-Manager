package api

import (
	"context"

	"github.com/google/go-github/v60/github"
	"golang.org/x/oauth2"
)

func CreateGithubClient(token string) *github.Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	tc := oauth2.NewClient(ctx, ts)

	return github.NewClient(tc)
}

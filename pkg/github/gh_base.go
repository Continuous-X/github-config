package github

import (
	"context"
	"github.com/google/go-github/v47/github"
	"golang.org/x/oauth2"
)

type GHBase struct {
	ghToken string
}

func (ghBase GHBase) getCient() (*github.Client, context.Context) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: ghBase.ghToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	return github.NewClient(tc), ctx
}

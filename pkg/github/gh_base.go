package github

import (
	"context"
	"log"
	"net/http"

	"github.com/google/go-github/v48/github"
	"golang.org/x/oauth2"
)

type GHBase struct {
	ghToken   string
	gheDomain string
}

func (ghBase GHBase) getCient() (*github.Client, context.Context) {
	if len(ghBase.gheDomain) > 0 {
		client, ctx, clientErr := ghBase.getEnterpriseCient()
		if clientErr != nil {
			log.Printf("error ghe client creation - %s \n switch to public gh client", clientErr)
		} else {
			return client, ctx
		}
	}
	client, ctx := ghBase.getPublicCient()
	return client, ctx
}

func (ghBase GHBase) getPublicCient() (*github.Client, context.Context) {
	tc, ctx := ghBase.getOAuthClient()
	return github.NewClient(tc), ctx
}

func (ghBase GHBase) getEnterpriseCient() (*github.Client, context.Context, error) {
	tc, ctx := ghBase.getOAuthClient()
	ghClient, ghClientError := github.NewEnterpriseClient("https://"+ghBase.gheDomain+"/api/v3/", "https://uploads."+ghBase.gheDomain+"/", tc)
	if ghClientError != nil {
		log.Printf("error ghe client creation - %s", ghClientError)
		return nil, ctx, ghClientError
	}
	return ghClient, ctx, ghClientError
}

func (ghBase GHBase) getOAuthClient() (*http.Client, context.Context) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: ghBase.ghToken},
	)
	return oauth2.NewClient(ctx, ts), ctx
}

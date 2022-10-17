package github

import (
	"fmt"
	"github-config/pkg/output"

	"github.com/google/go-github/v47/github"
)

type GHOrganization struct {
	Organisation   string
	GhToken        string
}

func (ghOrga GHOrganization) GetConfig() ([]*github.Organization, error) {

	client, ctx := GHBase{ghToken: ghOrga.GhToken}.getCient()
	organizationList, _, listError := client.Organizations.List(ctx, "", nil)
	if listError != nil {
		return organizationList, listError
	}
	for index, organization := range organizationList {
		output.PrintCliInfo(fmt.Sprintf("%d: %s", index, organization.GetLogin()))
	}
	return organizationList, nil
}

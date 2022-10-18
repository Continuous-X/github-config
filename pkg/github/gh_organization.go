package github

import (
	"encoding/json"
	error2 "github-config/pkg/error"
)

type GHOrganization struct {
	Organisation string
	GhToken      string
}

func (ghOrga GHOrganization) GetConfig(org string) (string, error) {

	client, ctx := GHBase{ghToken: ghOrga.GhToken}.getCient()
	organization, _, listError := client.Organizations.Get(ctx, org)
	if listError != nil {
		return "", listError
	}

	marshalledOrganization, err := json.Marshal(organization)
	if err != nil {
		error2.FailHandleCommand(err)
	}

	return string(marshalledOrganization), nil
}

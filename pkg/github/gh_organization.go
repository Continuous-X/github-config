package github

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

type GHOrganization struct {
	Organisation       string
	GhToken            string
	GhEnterpriseDomain string
}

func (ghOrga GHOrganization) GetConfig() (string, error) {

	client, ctx := GHBase{
		ghToken:   ghOrga.GhToken,
		gheDomain: ghOrga.GhEnterpriseDomain,
	}.getCient()
	organization, _, listError := client.Organizations.Get(ctx, ghOrga.Organisation)
	if listError != nil {
		return "", listError
	}

	yamlMarshal, yamlMarshalErr := yaml.Marshal(organization)
	if yamlMarshalErr != nil {
		return "", yamlMarshalErr
	}

	fmt.Printf("output:\n%s\n", yamlMarshal)

	return string(yamlMarshal), nil
}

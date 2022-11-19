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

	profile := &Profile{}
	yaml.Unmarshal(yamlMarshal,profile)
	profileMarshal, profileMarshalErr := yaml.Marshal(profile)
	if profileMarshalErr != nil {
		return "", profileMarshalErr
	}


	fmt.Printf("output:\n%s\n", profileMarshal)

	return string(profileMarshal), nil
}

type GHOrganizationSettings struct {
	Profile Profile `json:"profile"`
}

type Profile struct {
	Login                       *string    `json:"login,omitempty"`
	Name                        *string    `json:"name,omitempty"`
	Company                     *string    `json:"company,omitempty"`
	Blog                        *string    `json:"blog,omitempty"`
	Location                    *string    `json:"location,omitempty"`
	Email                       *string    `json:"email,omitempty"`
	TwitterUsername             *string    `json:"twitter_username,omitempty"`
	Description                 *string    `json:"description,omitempty"`
	BillingEmail                *string    `json:"billing_email,omitempty"`
	Type                        *string    `json:"type,omitempty"`
	TwoFactorRequirementEnabled *bool      `json:"two_factor_requirement_enabled,omitempty"`
	IsVerified                  *bool      `json:"is_verified,omitempty"`
	HasOrganizationProjects     *bool      `json:"has_organization_projects,omitempty"`
	HasRepositoryProjects       *bool      `json:"has_repository_projects,omitempty"`
}

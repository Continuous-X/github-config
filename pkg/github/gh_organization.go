package github

import (
	"github.com/google/go-github/v47/github"

	"gopkg.in/yaml.v2"
)

type GHOrganization struct {
	Organisation       string
	GhToken            string
	GhEnterpriseDomain string
}

func (ghOrga GHOrganization) GetRepositories() ([]*github.Repository, error) {
	client, ctx := GHBase{
		ghToken:   ghOrga.GhToken,
		gheDomain: ghOrga.GhEnterpriseDomain,
	}.getCient()
	var repos []*github.Repository
	var opts = github.RepositoryListByOrgOptions{}
	opts.PerPage = 50
	opts.Page = 1

	for opts.Page > 0 {
		readedRepoList, response, listError := client.Repositories.ListByOrg(ctx, ghOrga.Organisation, &opts)
		if listError != nil {
			return nil, listError
		}
		repos = append(repos, readedRepoList...)
		opts.Page = response.NextPage
	}

	return repos, nil
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

	profile := &GHOrgaProfile{}
	yaml.Unmarshal(yamlMarshal, profile)
	profileMarshal, profileMarshalErr := yaml.Marshal(profile)
	if profileMarshalErr != nil {
		return "", profileMarshalErr
	}

	return string(profileMarshal), nil
}

type GHOrganizationSettings struct {
	Profile GHOrgaProfile `json:"profile"`
}

type GHOrgaProfile struct {
	Login                       *string `json:"login,omitempty"`
	Name                        *string `json:"name,omitempty"`
	Company                     *string `json:"company,omitempty"`
	Blog                        *string `json:"blog,omitempty"`
	Location                    *string `json:"location,omitempty"`
	Email                       *string `json:"email,omitempty"`
	TwitterUsername             *string `json:"twitter_username,omitempty"`
	Description                 *string `json:"description,omitempty"`
	BillingEmail                *string `json:"billing_email,omitempty"`
	Type                        *string `json:"type,omitempty"`
	TwoFactorRequirementEnabled *bool   `json:"two_factor_requirement_enabled,omitempty"`
	IsVerified                  *bool   `json:"is_verified,omitempty"`
	HasOrganizationProjects     *bool   `json:"has_organization_projects,omitempty"`
	HasRepositoryProjects       *bool   `json:"has_repository_projects,omitempty"`
}

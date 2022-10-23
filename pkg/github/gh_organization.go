package github

import (
	"fmt"
	"time"

	"gopkg.in/yaml.v2"
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

	yamlMarshal, yamlMarshalErr := yaml.Marshal(organization)
	if yamlMarshalErr != nil {
		return "", yamlMarshalErr
	}

	fmt.Printf("output:\n%s\n", yamlMarshal)

	return string(yamlMarshal), nil
}

type GHOrganizationBackup struct {
	Login             string    `yaml:"login"`
	ID                int       `yaml:"id"`
	NodeID            string    `yaml:"node_id"`
	AvatarURL         string    `yaml:"avatar_url"`
	HTMLURL           string    `yaml:"html_url"`
	Name              string    `yaml:"name"`
	Description       string    `yaml:"description"`
	PublicRepos       int       `yaml:"public_repos"`
	PublicGists       int       `yaml:"public_gists"`
	Followers         int       `yaml:"followers"`
	Following         int       `yaml:"following"`
	CreatedAt         time.Time `yaml:"created_at"`
	UpdatedAt         time.Time `yaml:"updated_at"`
	TotalPrivateRepos int       `yaml:"total_private_repos"`
	OwnedPrivateRepos int       `yaml:"owned_private_repos"`
	PrivateGists      int       `yaml:"private_gists"`
	DiskUsage         int       `yaml:"disk_usage"`
	Collaborators     int       `yaml:"collaborators"`
	BillingEmail      string    `yaml:"billing_email"`
	Type              string    `yaml:"type"`
	Plan              struct {
		Name         string `yaml:"name"`
		Space        int    `yaml:"space"`
		PrivateRepos int    `yaml:"private_repos"`
		FilledSeats  int    `yaml:"filled_seats"`
		Seats        int    `yaml:"seats"`
	} `yaml:"plan"`
	TwoFactorRequirementEnabled                           bool   `yaml:"two_factor_requirement_enabled"`
	IsVerified                                            bool   `yaml:"is_verified"`
	HasOrganizationProjects                               bool   `yaml:"has_organization_projects"`
	HasRepositoryProjects                                 bool   `yaml:"has_repository_projects"`
	DefaultRepositoryPermission                           string `yaml:"default_repository_permission"`
	MembersCanCreateRepositories                          bool   `yaml:"members_can_create_repositories"`
	MembersCanCreatePublicRepositories                    bool   `yaml:"members_can_create_public_repositories"`
	MembersCanCreatePrivateRepositories                   bool   `yaml:"members_can_create_private_repositories"`
	MembersCanCreateInternalRepositories                  bool   `yaml:"members_can_create_internal_repositories"`
	MembersCanForkPrivateRepositories                     bool   `yaml:"members_can_fork_private_repositories"`
	MembersAllowedRepositoryCreationType                  string `yaml:"members_allowed_repository_creation_type"`
	MembersCanCreatePages                                 bool   `yaml:"members_can_create_pages"`
	MembersCanCreatePublicPages                           bool   `yaml:"members_can_create_public_pages"`
	MembersCanCreatePrivatePages                          bool   `yaml:"members_can_create_private_pages"`
	WebCommitSignoffRequired                              bool   `yaml:"web_commit_signoff_required"`
	AdvancedSecurityEnabledForNewRepositories             bool   `yaml:"advanced_security_enabled_for_new_repositories"`
	DependabotAlertsEnabledForNewRepositories             bool   `yaml:"dependabot_alerts_enabled_for_new_repositories"`
	DependabotSecurityUpdatesEnabledForNewRepositories    bool   `yaml:"dependabot_security_updates_enabled_for_new_repositories"`
	DependencyGraphEnabledForNewRepositories              bool   `yaml:"dependency_graph_enabled_for_new_repositories"`
	SecretScanningEnabledForNewRepositories               bool   `yaml:"secret_scanning_enabled_for_new_repositories"`
	SecretScanningPushProtectionEnabledForNewRepositories bool   `yaml:"secret_scanning_push_protection_enabled_for_new_repositories"`
	URL                                                   string `yaml:"url"`
	EventsURL                                             string `yaml:"events_url"`
	HooksURL                                              string `yaml:"hooks_url"`
	IssuesURL                                             string `yaml:"issues_url"`
	MembersURL                                            string `yaml:"members_url"`
	PublicMembersURL                                      string `yaml:"public_members_url"`
	ReposURL                                              string `yaml:"repos_url"`
}

package cmd

import (
	"fmt"
	"github-config/pkg/github"
	"github-config/pkg/output"

	"github.com/spf13/cobra"
)

var organizationCmd = &cobra.Command{
	Use:   "organization",
	Short: "export the github organization config",
	Long: `export the github organization config. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		LogOutput.AddLoggingLine(output.LogTypeInfo, cmd.Name(), "command called")

		fmt.Println("export called")
		gh_personal_token, _ := cmd.Flags().GetString(flag_gh_token)
		output.PrintCliInfo(fmt.Sprintf("%s - '%s'", flag_gh_token, gh_personal_token))

		gh_organization, _ := cmd.Flags().GetString(flag_gh_orga)
		output.PrintCliInfo(fmt.Sprintf("%s - '%s'", flag_gh_orga, gh_organization))

		gh_repository, _ := cmd.Flags().GetString(flag_gh_repo)
		output.PrintCliInfo(fmt.Sprintf("%s - '%s'", flag_gh_repo, gh_repository))

		if len(gh_repository) > 0 {
			repoConfig, repoConfigErr := github.GHRepository{
				Organisation: github.GHOrganization{
					Organisation:       gh_organization,
					GhToken:            gh_personal_token,
					GhEnterpriseDomain: Config.Github.EnterpriseDomain,
				},
				Repository: gh_repository,
			}.GetConfig()
			if repoConfigErr != nil {
				LogOutput.AddLoggingLine(output.LogTypeError, "export", repoConfigErr.Error())
			} else {
				github.GHRepositoryContent{
					Organisation:       Config.Export.Github.Organization,
					RepositoryName:     Config.Export.Github.Repository,
					GhToken:            Config.Export.Github.Token,
					GhEnterpriseDomain: Config.Github.EnterpriseDomain,
				}.WriteContent(
					fmt.Sprintf("orgs/%s/repos/%s/repository-config.yaml", gh_organization, gh_repository),
					"main",
					repoConfig,
					fmt.Sprintf("export config from github repository '%s'", gh_repository),
					"Lyle",
					"lyle@github.com",
				)
			}
		} else {
			orgaConfig, orgaConfigErr := github.GHOrganization{
				Organisation:       gh_organization,
				GhToken:            gh_personal_token,
				GhEnterpriseDomain: Config.Github.EnterpriseDomain,
			}.GetConfig()
			if orgaConfigErr != nil {
				LogOutput.AddLoggingLine(output.LogTypeError, "export", orgaConfigErr.Error())
			} else {
				github.GHRepositoryContent{
					Organisation:       Config.Export.Github.Organization,
					RepositoryName:     Config.Export.Github.Repository,
					GhToken:            Config.Export.Github.Token,
					GhEnterpriseDomain: Config.Github.EnterpriseDomain,
				}.WriteContent(
					fmt.Sprintf("orgs/%s/organization-config.yaml", gh_organization),
					"main",
					orgaConfig,
					fmt.Sprintf("export config from github organization '%s'", gh_organization),
					"Lyle",
					"lyle@github.com",
				)
			}
		}

		LogOutput.PrintLogging()

	},
}

func init() {
	exportCmd.AddCommand(organizationCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// exportCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// exportCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	organizationCmd.Flags().StringVarP(&gh_personal_token, flag_gh_token, flag_gh_token_short, "", flag_gh_token_description)
	organizationCmd.Flags().StringVarP(&gh_organization, flag_gh_orga, flag_gh_orga_short, "", flag_gh_orga_description)
	organizationCmd.Flags().StringVarP(&gh_repository, flag_gh_repo, flag_gh_repo_short, "", flag_gh_repo_description)
	organizationCmd.MarkFlagRequired(flag_gh_token)
	organizationCmd.MarkFlagRequired(flag_gh_orga)
}

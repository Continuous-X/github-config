package cmd

import (
	"fmt"
	"github-config/pkg/github"
	"github-config/pkg/output"

	"github.com/spf13/cobra"
)

var organizationCmd = &cobra.Command{
	Use:   cmd_organization,
	Short: "export the github organization config",
	Long: `export the github organization configuration in the backup repository.

	......`,
	Run: func(cmd *cobra.Command, args []string) {
		LogOutput.AddLoggingLine(output.LogTypeInfo, cmd.CommandPath(), "command started")

		readed_flag_gh_token := cmd.Flag(flag_gh_token)
		if readed_flag_gh_token.Changed {
			// TODO: input validation
			gh_personal_token = readed_flag_gh_token.Value.String()
			output.PrintCliInfo(fmt.Sprintf("%s - '%s'", flag_gh_token, gh_personal_token))
		}

		readed_flag_gh_orga := cmd.Flag(flag_gh_orga)
		if readed_flag_gh_orga.Changed {
			// TODO: input validation
			gh_organization = readed_flag_gh_orga.Value.String()
			output.PrintCliInfo(fmt.Sprintf("%s - '%s'", flag_gh_orga, gh_organization))
		}

		if readed_flag_all_gh_repos := cmd.Flag(flag_all_gh_repos); readed_flag_all_gh_repos.Changed {
			output.PrintCliInfo(fmt.Sprintf("%s - '%v'", flag_all_gh_repos, readed_flag_all_gh_repos.Changed))
			LogOutput.AddLoggingLine(output.LogTypeInfo, cmd.CommandPath(), "export all repository and organization configuration")
		} else {
			output.PrintCliInfo(fmt.Sprintf("%s - '%s'", flag_all_gh_repos, "export organization configuration"))
		}

		orgaConfig, orgaConfigErr := github.GHOrganization{
			Organisation:       gh_organization,
			GhToken:            gh_personal_token,
			GhEnterpriseDomain: Config.Github.EnterpriseDomain,
		}.GetConfig()
		if orgaConfigErr != nil {
			LogOutput.AddLoggingLine(output.LogTypeError, cmd.CommandPath(), orgaConfigErr.Error())
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

		LogOutput.AddLoggingLine(output.LogTypeInfo, cmd.CommandPath(), "command ended")
		LogOutput.PrintLogging()

	},
}

func init() {
	exportCmd.AddCommand(organizationCmd)

	organizationCmd.Flags().StringVarP(&gh_personal_token, flag_gh_token, flag_gh_token_short, "", flag_gh_token_description)
	organizationCmd.Flags().StringVarP(&gh_organization, flag_gh_orga, flag_gh_orga_short, "", flag_gh_orga_description)
	organizationCmd.Flags().BoolP(flag_all_gh_repos, flag_all_gh_repos_short, false, flag_all_gh_repos_description)

	organizationCmd.MarkFlagRequired(flag_gh_token)
	organizationCmd.MarkFlagRequired(flag_gh_orga)
}

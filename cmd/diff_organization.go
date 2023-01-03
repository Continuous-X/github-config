package cmd

import (
	"fmt"
	"github-config/pkg/github"
	"reflect"

	"gopkg.in/yaml.v2"

	"golang.org/x/exp/slog"

	"github.com/spf13/cobra"
)

var diffOrganizationCmd = &cobra.Command{
	Use:   cmd_organization,
	Short: "export the github organization config",
	Long: `export the github organization configuration in the backup repository.

	......`,
	Run: func(cmd *cobra.Command, args []string) {
		slog.Debug("command started", "cmd", cmd.CommandPath())

		readed_flag_gh_token := cmd.Flag(flag_gh_token)
		if readed_flag_gh_token.Changed {
			// TODO: input validation
			gh_personal_token = readed_flag_gh_token.Value.String()
			slog.Info(fmt.Sprintf("%s - '%s'", flag_gh_token, gh_personal_token))
		}

		readed_flag_gh_orga := cmd.Flag(flag_gh_orga)
		if readed_flag_gh_orga.Changed {
			// TODO: input validation
			gh_organization = readed_flag_gh_orga.Value.String()
			slog.Info(fmt.Sprintf("%s - '%s'", flag_gh_orga, gh_organization))
		}

		orgaConfig, orgaConfigErr := github.GHOrganization{
			Organisation:       gh_organization,
			GhToken:            gh_personal_token,
			GhEnterpriseDomain: Config.Github.EnterpriseDomain,
		}.GetConfig()
		if orgaConfigErr != nil {
			slog.Error("oops", orgaConfigErr)
		} else {
			fileContent, fileContentErr := github.GHRepositoryContent{
				Organisation:       Config.Export.Github.Organization,
				RepositoryName:     Config.Export.Github.Repository,
				GhToken:            Config.Export.Github.Token,
				GhEnterpriseDomain: Config.Github.EnterpriseDomain,
			}.GetFileContentDecoded(
				fmt.Sprintf("orgs/%s/organization-config.yaml", gh_organization),
				"main",
			)
			if fileContentErr != nil {
				slog.Error("oops", fileContentErr)
			}
			profileFromBackup := &github.GHOrgaProfile{}
			profileFromReadedOrga := &github.GHOrgaProfile{}
			yaml.Unmarshal([]byte(fileContent), profileFromBackup)
			slog.Info("readed from backup file", profileFromBackup)
			yaml.Unmarshal([]byte(orgaConfig), profileFromReadedOrga)
			slog.Info("readed from gh orga", profileFromReadedOrga)

			reflect.DeepEqual(profileFromBackup, profileFromReadedOrga)

		}

		slog.Debug("command ended", "cmd", cmd.CommandPath())
	},
}

func init() {
	diffCmd.AddCommand(diffOrganizationCmd)

	diffOrganizationCmd.Flags().StringVarP(&gh_personal_token, flag_gh_token, flag_gh_token_short, "", flag_gh_token_description)
	diffOrganizationCmd.Flags().StringVarP(&gh_organization, flag_gh_orga, flag_gh_orga_short, "", flag_gh_orga_description)

	diffOrganizationCmd.MarkFlagRequired(flag_gh_token)
	diffOrganizationCmd.MarkFlagRequired(flag_gh_orga)
}

package cmd

import (
	"fmt"
	"github-config/pkg/github"
	"os"

	"gopkg.in/yaml.v2"

	"golang.org/x/exp/slog"

	"github.com/spf13/cobra"
)

var syncOrganizationCmd = &cobra.Command{
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
		yaml.Unmarshal([]byte(fileContent), profileFromBackup)

		orgaWriteErr := github.GHOrganization{
			Organisation:       gh_organization,
			GhToken:            gh_personal_token,
			GhEnterpriseDomain: Config.Github.EnterpriseDomain,
		}.WriteConfigProfile(*profileFromBackup)

		if orgaWriteErr != nil {
			slog.Error("oops", orgaWriteErr)
			os.Exit(1)
		} else {
			slog.Info("updated from backup file", profileFromBackup)
		}

		slog.Debug("command ended", "cmd", cmd.CommandPath())
	},
}

func init() {
	syncCmd.AddCommand(syncOrganizationCmd)

	syncOrganizationCmd.Flags().StringVarP(&gh_personal_token, flag_gh_token, flag_gh_token_short, "", flag_gh_token_description)
	syncOrganizationCmd.Flags().StringVarP(&gh_organization, flag_gh_orga, flag_gh_orga_short, "", flag_gh_orga_description)

	syncOrganizationCmd.MarkFlagRequired(flag_gh_token)
	syncOrganizationCmd.MarkFlagRequired(flag_gh_orga)
}

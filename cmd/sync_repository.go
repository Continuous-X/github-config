package cmd

import (
	"fmt"
	"github-config/pkg/github"
	"os"

	"gopkg.in/yaml.v2"

	"golang.org/x/exp/slog"

	"github.com/spf13/cobra"
)

var syncRepositoryCmd = &cobra.Command{
	Use:   cmd_sync,
	Short: "sync the github organization config",
	Long: `sync the github organization configuration in the backup repository.
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

		readed_flag_gh_repo := cmd.Flag(flag_gh_repo)
		if readed_flag_gh_repo.Changed {
			// TODO: input validation
			gh_repository = readed_flag_gh_repo.Value.String()
			slog.Info(fmt.Sprintf("%s - '%s'", flag_gh_repo, gh_repository))
		}


		fileContent, fileContentErr := github.GHRepositoryContent{
			Organisation:       Config.Export.Github.Organization,
			RepositoryName:     Config.Export.Github.Repository,
			GhToken:            Config.Export.Github.Token,
			GhEnterpriseDomain: Config.Github.EnterpriseDomain,
		}.GetFileContentDecoded(
			fmt.Sprintf("orgs/%s/repos/%s/repository-config.yaml", gh_organization, gh_repository),
			"main",
		)
		if fileContentErr != nil {
			slog.Error("oops", fileContentErr)
		}
		profileFromBackup := &github.GHRepositoryProfile{}
		yaml.Unmarshal([]byte(fileContent), profileFromBackup)

		repoWriteErr := github.GHRepository{
			Organisation: github.GHOrganization{
				Organisation:       gh_organization,
				GhToken:            gh_personal_token,
				GhEnterpriseDomain: Config.Github.EnterpriseDomain,
			},
			Repository: gh_repository,
		}.WriteConfigProfile(*profileFromBackup)

		if repoWriteErr != nil {
			slog.Error("oops", repoWriteErr)
			os.Exit(1)
		} else {
			slog.Info("updated from backup file", profileFromBackup)
		}

		slog.Debug("command ended", "cmd", cmd.CommandPath())
	},
}

func init() {
	repositoryCmd.AddCommand(syncRepositoryCmd)

	syncRepositoryCmd.Flags().StringVarP(&gh_personal_token, flag_gh_token, flag_gh_token_short, "", flag_gh_token_description)
	syncRepositoryCmd.Flags().StringVarP(&gh_organization, flag_gh_orga, flag_gh_orga_short, "", flag_gh_orga_description)
	syncRepositoryCmd.Flags().StringVarP(&gh_repository, flag_gh_repo, flag_gh_repo_short, "", flag_gh_repo_description)

	syncRepositoryCmd.MarkFlagRequired(flag_gh_token)
	syncRepositoryCmd.MarkFlagRequired(flag_gh_orga)
	syncRepositoryCmd.MarkFlagRequired(flag_gh_repo)

}

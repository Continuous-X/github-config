/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github-config/pkg/github"

	"golang.org/x/exp/slog"

	"github.com/spf13/cobra"
)

// exportRepositoryCmd represents the export command
var exportRepositoryCmd = &cobra.Command{
	Use:   cmd_export,
	Short: "export repository ",
	Long:  `export repository data`,
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

		repoConfig, repoConfigErr := github.GHRepository{
			Organisation: github.GHOrganization{
				Organisation:       gh_organization,
				GhToken:            gh_personal_token,
				GhEnterpriseDomain: Config.Github.EnterpriseDomain,
			},
			Repository: gh_repository,
		}.GetConfig()

		if repoConfigErr != nil {
			slog.Error("oops", repoConfigErr)
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

		slog.Debug("command ended", "cmd", cmd.CommandPath())
	},
}

func init() {
	repositoryCmd.AddCommand(exportRepositoryCmd)

	exportRepositoryCmd.Flags().StringVarP(&gh_personal_token, flag_gh_token, flag_gh_token_short, "", flag_gh_token_description)
	exportRepositoryCmd.Flags().StringVarP(&gh_organization, flag_gh_orga, flag_gh_orga_short, "", flag_gh_orga_description)
	exportRepositoryCmd.Flags().StringVarP(&gh_repository, flag_gh_repo, flag_gh_repo_short, "", flag_gh_repo_description)

	exportRepositoryCmd.MarkFlagRequired(flag_gh_token)
	exportRepositoryCmd.MarkFlagRequired(flag_gh_orga)
	exportRepositoryCmd.MarkFlagRequired(flag_gh_repo)

}

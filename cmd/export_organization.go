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

var exportOrganizationCmd = &cobra.Command{
	Use:   cmd_export,
	Short: "export organization ",
	Long:  `export organization data`,
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

		readed_flag_all_gh_repos := cmd.Flag(flag_all_gh_repos)
		slog.Debug(fmt.Sprintf("%s - '%v'", flag_all_gh_repos, readed_flag_all_gh_repos.Changed))

		orgaConfig, orgaConfigErr := github.GHOrganization{
			Organisation:       gh_organization,
			GhToken:            gh_personal_token,
			GhEnterpriseDomain: Config.Github.EnterpriseDomain,
		}.GetConfig()
		if orgaConfigErr != nil {
			slog.Error("oops", orgaConfigErr)
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

		if readed_flag_all_gh_repos.Changed {
			slog.Info("export all repository and organization configuration", "cmd", cmd.CommandPath())
			repoList, repoListErr := github.GHOrganization{
				Organisation:       gh_organization,
				GhToken:            gh_personal_token,
				GhEnterpriseDomain: Config.Github.EnterpriseDomain,
			}.GetRepositories()
			if repoListErr != nil {
				slog.Error("oops", repoListErr)
			} else {
				for _, repo := range repoList {
					repoConfig, repoConfigErr := github.GHRepository{
						Organisation: github.GHOrganization{
							Organisation:       gh_organization,
							GhToken:            gh_personal_token,
							GhEnterpriseDomain: Config.Github.EnterpriseDomain,
						},
						Repository: *repo.Name,
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
							fmt.Sprintf("orgs/%s/repos/%s/repository-config.yaml", gh_organization, *repo.Name),
							"main",
							repoConfig,
							fmt.Sprintf("export config from github repository '%s'", *repo.Name),
							"Lyle",
							"lyle@github.com",
						)
					}
				}
			}
		} else {
			slog.Info("export just organization configuration", "cmd", cmd.CommandPath())
		}

		slog.Debug("command ended", "cmd", cmd.CommandPath())
	},
}

func init() {
	organizationCmd.AddCommand(exportOrganizationCmd)

	exportOrganizationCmd.Flags().StringVarP(&gh_personal_token, flag_gh_token, flag_gh_token_short, "", flag_gh_token_description)
	exportOrganizationCmd.Flags().StringVarP(&gh_organization, flag_gh_orga, flag_gh_orga_short, "", flag_gh_orga_description)
	exportOrganizationCmd.Flags().BoolP(flag_all_gh_repos, flag_all_gh_repos_short, false, flag_all_gh_repos_description)

	exportOrganizationCmd.MarkFlagRequired(flag_gh_token)
	exportOrganizationCmd.MarkFlagRequired(flag_gh_orga)
}

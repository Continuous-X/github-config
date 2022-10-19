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
	"github-config/pkg/output"

	"github.com/spf13/cobra"
)

var (
	gh_personal_token string
	gh_organization   string
)

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("export called")
		gh_p_flag, _ := cmd.Flags().GetString(flag_gh_token)
		output.PrintCliInfo(fmt.Sprintf("%s - '%s'", flag_gh_token,gh_p_flag))

		gh_o_flag, _ := cmd.Flags().GetString(flag_gh_orga)
		output.PrintCliInfo(fmt.Sprintf("%s - '%s'", flag_gh_orga, gh_o_flag))

		fmt.Println(github.GHOrganization{Organisation: "Continuous-X", GhToken: gh_p_flag}.GetConfig(gh_o_flag))

		fmt.Printf("github domain: %s\ngithub organization: %s\ngithub repository: %s\n", GithubDomain, GithubOrga, GithubRepo)

	},
}

func init() {
	rootCmd.AddCommand(exportCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// exportCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// exportCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	exportCmd.Flags().StringVarP(&gh_personal_token, flag_gh_token, flag_gh_token_short, "", flag_gh_token_description)
	exportCmd.Flags().StringVarP(&gh_organization, flag_gh_orga, flag_gh_orga_short, "", flag_gh_orga_description)
	exportCmd.MarkFlagRequired(flag_gh_token)
	exportCmd.MarkFlagRequired(flag_gh_orga)
}

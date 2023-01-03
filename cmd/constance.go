package cmd

const (
	flag_gh_token                 = "token"
	flag_gh_token_short           = "t"
	flag_gh_token_description     = "github personal token"
	flag_gh_orga                  = "organization"
	flag_gh_orga_short            = "o"
	flag_gh_orga_description      = "github organization"
	flag_gh_repo                  = "repository"
	flag_gh_repo_short            = "r"
	flag_gh_repo_description      = "github repository"
	flag_all_gh_repos             = "all-repositories"
	flag_all_gh_repos_short       = ""
	flag_all_gh_repos_description = "[default false] / all github repositories in organization"
	flag_dry_run                  = "dry-run"
	flag_dry_run_short            = ""
	flag_dry_run_default          = "none"
	flag_dry_run_description      = "[default none] / Must be \"none\" or \"server\". If server strategy, only print the config informations that would be sent"

	// commands
	cmd_export       = "export"
	cmd_diff         = "diff"
	cmd_sync         = "sync"
	cmd_organization = "organization"
	cmd_repository   = "repository"
)

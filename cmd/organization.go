package cmd

import (
	"golang.org/x/exp/slog"

	"github.com/spf13/cobra"
)

var organizationCmd = &cobra.Command{
	Use:   cmd_organization,
	Short: cmd_organization_desc_short,
	Long:  cmd_organization_desc_long,
	Run: func(cmd *cobra.Command, args []string) {
		slog.Debug("command started", "cmd", cmd.CommandPath())
		cmd.Usage()
		slog.Debug("command ended", "cmd", cmd.CommandPath())

	},
}

func init() {
	rootCmd.AddCommand(organizationCmd)
}

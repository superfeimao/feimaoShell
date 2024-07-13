package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:     "shell",
	Aliases: []string{"feimaoShell"},
	Short:   "feima",
	Long:    `Qt-core-shell is a CLI tool controls Qingteng MicroStep`,
	Run: func(cmd *cobra.Command, args []string) {
		shell.Run()
	},
}

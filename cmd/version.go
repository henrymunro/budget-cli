package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = "0.1.0"

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of budget-cli",
	Long:  `All software has versions. This is budget-cli's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version)
	},
}

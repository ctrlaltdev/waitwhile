package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version",
	Long:  `Print the version of this WaitWhile CLI executable`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("WaitWhile CLI - %s\n", rootCmd.Version)
	},
}

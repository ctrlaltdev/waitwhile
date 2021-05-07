package cmd

import (
	"fmt"

	"github.com/ctrlaltdev/waitwhile/waitwhile"

	"github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
)

var (
	version  = "v1.0.0"
	apiKey   string
	configDB = mysql.Config{Net: "tcp", AllowNativePasswords: true}
)

var rootCmd = &cobra.Command{
	Version: version,
	Use:     "waitwhile",
	Aliases: []string{"ww"},
	Short:   "WaitWhile CLI",
	Long:    `This is a CLI Tool to execute some operations on WaitWhile using their APIs`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		waitwhile.Init(&apiKey)
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.SetVersionTemplate(fmt.Sprintf("WaitWhile CLI - %s\n", rootCmd.Version))

	rootCmd.PersistentFlags().StringVarP(&apiKey, "apikey", "a", "", "Center API Key (required)")
	rootCmd.MarkPersistentFlagRequired("apikey")

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(locationsCmd)
}

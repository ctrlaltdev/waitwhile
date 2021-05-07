package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var locationsDeleteCmd = &cobra.Command{
	Use:     "delete [location id (optional)]",
	Short:   "Delete Locations",
	Long:    `Delete a Location or all the Locations for that center`,
	Aliases: []string{"rm", "del"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			fmt.Printf("Delete %s", args[0])
		} else {
			fmt.Printf("Delete All")
		}
	},
}

func init() {

}

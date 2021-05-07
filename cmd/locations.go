package cmd

import (
	"github.com/spf13/cobra"
)

var locationsCmd = &cobra.Command{
	Use:   "locations",
	Short: "Interact with Locations",
	Long:  `Interact with Locations Entities`,
}

func init() {
	locationsCmd.AddCommand(locationsGetCmd)
	locationsCmd.AddCommand(locationsSaveCmd)
	locationsCmd.AddCommand(locationsUpdateCmd)
	locationsCmd.AddCommand(locationsDeleteCmd)
}

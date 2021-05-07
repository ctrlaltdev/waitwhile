package cmd

import (
	"github.com/ctrlaltdev/waitwhile/utils"
	"github.com/ctrlaltdev/waitwhile/waitwhile"

	"github.com/spf13/cobra"
)

var (
	payload string
)

var locationsUpdateCmd = &cobra.Command{
	Use:     "update [location id (optional)]",
	Short:   "Update Locations",
	Long:    `Update a Location or all the Locations for that center`,
	Aliases: []string{"read", "ls"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			utils.CheckErr(waitwhile.UpdateLocation(&args[0], &payload))
		} else {
			var locations []waitwhile.Location

			var startAfter string
			limit := 100

			for startAfter != "END" {
				body := waitwhile.GetLocations(limit, &startAfter)
				startAfter = body.EndAt
				locations = append(locations, body.Results...)

				if startAfter == "" {
					startAfter = "END"
				}
			}
			for _, location := range locations {
				utils.CheckErr(waitwhile.UpdateLocation(&location.ID, &payload))
			}
		}
	},
}

func init() {
	locationsUpdateCmd.Flags().StringVarP(&payload, "payload", "p", "", "Payload for the update (required)")
	locationsUpdateCmd.MarkFlagRequired("payload")
}

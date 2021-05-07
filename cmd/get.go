package cmd

import (
	"fmt"

	"github.com/ctrlaltdev/waitwhile/waitwhile"

	"github.com/spf13/cobra"
)

var (
	locGetLimit int
)

var locationsGetCmd = &cobra.Command{
	Use:     "get [location id (optional)]",
	Short:   "Get Locations",
	Long:    `Get a Location or all the Locations for that center`,
	Aliases: []string{"read", "ls"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			location := waitwhile.GetLocation(&args[0])
			fmt.Println(location)
		} else {
			var locations []waitwhile.Location

			var startAfter string

			for startAfter != "END" {
				body := waitwhile.GetLocations(locGetLimit, &startAfter)
				startAfter = body.EndAt
				locations = append(locations, body.Results...)

				if startAfter == "" {
					startAfter = "END"
				}
			}
			for _, location := range locations {
				fmt.Println(location)
			}
		}
	},
}

func init() {
	locationsGetCmd.Flags().IntVarP(&locGetLimit, "limit", "l", 100, "Maximum items to retrieve per batch")
}

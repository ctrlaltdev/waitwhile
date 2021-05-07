package cmd

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/ctrlaltdev/waitwhile/utils"
	"github.com/ctrlaltdev/waitwhile/waitwhile"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
)

var (
	locSaveLimit int
)

var locationsSaveCmd = &cobra.Command{
	Use:     "save [location id (optional)]",
	Short:   "Save Locations",
	Long:    `Save a Location or all the Locations for that center`,
	Aliases: []string{"cp"},
	Run: func(cmd *cobra.Command, args []string) {
		db, err := sql.Open("mysql", configDB.FormatDSN())
		utils.CheckErr(err)

		defer db.Close()

		if len(args) > 0 {
			location := waitwhile.GetLocation(&args[0])
			fmt.Println(location)
		} else {
			var locations []waitwhile.Location

			var startAfter string

			for startAfter != "END" {
				body := waitwhile.GetLocations(locSaveLimit, &startAfter)
				startAfter = body.EndAt
				locations = append(locations, body.Results...)

				if startAfter == "" {
					startAfter = "END"
				}
			}
			for _, location := range locations {
				stmt, err := db.Prepare("INSERT INTO locations (BusinessName, BusinessType, Created, CreatedBy, Email, ID, IsActive, IsForceClosed, IsForceOpen, IsPublicBooking, IsPublicCheckIn, IsPublicWaitlist, Name, ShortName, Locale, Updated, UpdatedBy) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
				utils.CheckErr(err)

				_, err = stmt.Exec(location.BusinessName, location.BusinessType, location.Created, location.CreatedBy, location.Email, location.ID, location.IsActive, location.IsForceClosed, location.IsForceOpen, location.IsPublicBooking, location.IsPublicCheckIn, location.IsPublicWaitlist, location.Name, location.ShortName, location.Locale, location.Updated, location.UpdatedBy)
				utils.CheckErr(err)
			}
		}
	},
}

func init() {
	locationsSaveCmd.Flags().IntVarP(&locSaveLimit, "limit", "l", 100, "Maximum items to retrieve per batch")

	locationsSaveCmd.Flags().StringVar(&configDB.User, "dbuser", os.Getenv("DBUSER"), "Database User")
	locationsSaveCmd.Flags().StringVar(&configDB.Passwd, "dbpass", os.Getenv("DBPASS"), "Database Pass")
	locationsSaveCmd.Flags().StringVar(&configDB.Addr, "dbhost", "localhost", "Database Host")
	locationsSaveCmd.Flags().StringVar(&configDB.DBName, "dbname", "waitwhile", "Database Name")

	locationsSaveCmd.MarkFlagRequired("dbuser")
	locationsSaveCmd.MarkFlagRequired("dbpass")
}

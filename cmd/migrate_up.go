package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/aniketexplorex/dbmigrator/database"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
)

var migrateUpCmd *cobra.Command

func init() {
	migrateUpCmd = &cobra.Command{
		Use:   "up",
		Short: "migrate to v1 command",
		Long:  `Command to install version 1 of our application`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Running migrate up command")
			db := database.Open()
			dbDriver, err := mysql.WithInstance(db, &mysql.Config{})
			if err != nil {
				log.Fatalf("Error creating mySQL database instance %v", err)
				return
			}
			fileSource, err := (&file.File{}).Open(os.Getenv("MIGRATIONS_DIR"))
			if err != nil {
				log.Fatalf("Error opening file %v\n", err)
				return
			}
			m, err := migrate.NewWithInstance("file", fileSource, os.Getenv("DATABASE_NAME"), dbDriver)
			if err != nil {
				log.Fatalf("Migrate error %v\n", err)
				return
			}

			// if err := m.Up(); err != nil && err.Error() != "no change" {
			// 	log.Fatalf("Migrate up error %v\n", err)
			// 	return
			// }

			fmt.Println("Migration Up done")
		},
	}

	migrationCmd.AddCommand(migrateUpCmd)
}

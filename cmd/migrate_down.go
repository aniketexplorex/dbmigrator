package cmd

import (
	"fmt"
	"github.com/byteintellect/dbmigrator/database"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
	"os"
)

var migrateDownCmd *cobra.Command

func init() {
	migrateDownCmd = &cobra.Command{
		Use:   "down",
		Short: "migrate from v2 to v1 command",
		Long:  `Command to downgrade from v2 to v1`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Running migrate down command")
			db := database.Open()

			dbDriver, err := mysql.WithInstance(db, &mysql.Config{})
			if err != nil {
				fmt.Printf("Error creating instance %v\n", err)
				return
			}

			fileSource, err := (&file.File{}).Open(os.Getenv("MIGRATIONS_DIR"))
			if err != nil {
				fmt.Printf("Error opening migration %v\n", err)
				return
			}

			m, err := migrate.NewWithInstance("file", fileSource, os.Getenv("DATABASE_NAME"), dbDriver)
			if err != nil {
				fmt.Printf("Error creating migration instance %v\n", err)
				return
			}

			if err = m.Down(); err != nil {
				fmt.Printf("Error running down migration %v\n", err)
				return
			}

			fmt.Printf("Migrate down finished successfully\n")
		},
	}
	migrationCmd.AddCommand(migrateDownCmd)
}

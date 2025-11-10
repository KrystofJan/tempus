package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/KrystofJan/tempus/internal/handlers"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

var showEntryCmd = &cobra.Command{
	Use:   "show",
	Short: "Show entries",
	RunE: func(cmd *cobra.Command, args []string) error {
		all, err := cmd.Flags().GetBool("all")
		if err != nil {
			return fmt.Errorf("PARAMETER ERROR: %v", err)
		}

		if all {
			return handlers.ShowAllEntries()
		}

		id, err := cmd.Flags().GetInt64("id")
		if err != nil || id == 0 {
			return fmt.Errorf("PARAMETER ERROR: %v", err)
		}

		return handlers.ShowEntryById(id)
	},
}

func init() {
	showEntryCmd.Flags().BoolP("all", "a", false, "Show all")
	showEntryCmd.Flags().Int64P("id", "i", 0, "entry id")
	entryCmd.AddCommand(showEntryCmd)
}

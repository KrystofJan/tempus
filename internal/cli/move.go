package cli

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/KrystofJan/tempus/internal/handlers"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

var moveCmd = &cobra.Command{
	Use:   "move",
	Short: "move entry to the task",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			fmt.Println(args)
			return fmt.Errorf("This command needs 2 parameters:\n\t1. ID of the entry you want to move\n\t2. ID of the Task you want to move it to\nPlease provide the program with these values")
		}
		entryId, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("Could not convert the entry id to integer!\nOriginalError: %v", err)
		}
		taskId, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			return fmt.Errorf("Could not convert the task id to integer!\nOriginalError: %v", err)
		}

		if err := handlers.MoveEntry(entryId, taskId); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(moveCmd)
}

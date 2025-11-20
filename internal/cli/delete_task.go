package cli

import (
	"fmt"
	"strconv"

	"github.com/KrystofJan/tempus/internal/handlers"
	"github.com/spf13/cobra"
)

var deleteTaskCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete task",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("This command needs to have an ID of the to be deleted entry to be passes as a parameter")
		}

		taskId, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("Could not convert the entry id to integer!\nOriginalError: %v", err)
		}

		return handlers.DeleteTask(taskId)
	},
}

func init() {
	taskCmd.AddCommand(deleteTaskCmd)
}

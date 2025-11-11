package handlers

import (
	"github.com/KrystofJan/tempus/internal/service"
)

func MoveEntry(entryId, taskId int64) error {
	entryProvider, err := service.NewEntryProvider()
	if err != nil {
		return nil
	}

	return entryProvider.MoveTask(entryId, taskId)
}

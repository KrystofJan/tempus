package handlers

import (
	"github.com/KrystofJan/tempus/internal/db"
	"github.com/KrystofJan/tempus/internal/service"
)

func DeleteTask(id int64) error {
	db, err := db.NewDatabase()
	if err != nil {
		return err
	}
	taskProvider, err := service.NewTaskProvider(db)
	if err != nil {
		return nil
	}
	taskProvider.DeleteTask(id)
	return nil
}

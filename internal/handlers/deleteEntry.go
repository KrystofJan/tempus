package handlers

import (
	"github.com/KrystofJan/tempus/internal/db"
	"github.com/KrystofJan/tempus/internal/service"
)

func DeleteEntry(id int64) error {
	db, err := db.NewDatabase()
	if err != nil {
		return err
	}
	entryProvider, err := service.NewEntryProvider(db)
	if err != nil {
		return nil
	}
	return entryProvider.DeleteEntry(id)
}

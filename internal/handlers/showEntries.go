package handlers

import (
	"database/sql"
	"fmt"

	"github.com/KrystofJan/tempus/internal/display"
	"github.com/KrystofJan/tempus/internal/repository"
	"github.com/KrystofJan/tempus/internal/service"
)

func ShowEntryById(id int64) error {
	entryProvider, err := service.NewEntryProvider()
	if err != nil {
		return err
	}
	entry, err := entryProvider.FindEntryById(id)
	if err != nil {
		return fmt.Errorf("SERVICE ERROR: %v", err)
	}
	entries := []repository.Entry{entry}
	display.PrintEntries(entries)
	return nil
}

func ShowAllEntries() error {
	entryProvider, err := service.NewEntryProvider()
	if err != nil {
		return err
	}
	entries, err := entryProvider.FindAllEntries()
	if err != nil {
		return fmt.Errorf("SERVICE ERROR: %v", err)
	}
	for i := range len(entries) {
		recordedTime, err := entryProvider.CalculateEntryTime(entries[i].ID)
		if err != nil {
			return err
		}
		entries[i].RecordedTime = sql.NullInt64{Int64: recordedTime, Valid: true}
	}
	display.PrintEntries(entries)
	return nil
}

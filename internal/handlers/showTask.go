package handlers

import (
	"database/sql"
	"fmt"

	"github.com/KrystofJan/tempus/internal/db"
	"github.com/KrystofJan/tempus/internal/display"
	"github.com/KrystofJan/tempus/internal/repository"
	"github.com/KrystofJan/tempus/internal/service"
)

func ShowTaskById(id int64) error {
	db, err := db.NewDatabase()
	if err != nil {
		return err
	}
	taskProvider, err := service.NewTaskProvider(db)
	if err != nil {
		return err
	}
	task, err := taskProvider.FindTaskById(id)
	if err != nil {
		return fmt.Errorf("SERVICE ERROR: %v", err)
	}
	tasks := []repository.Task{task}
	display.PrintTasks(tasks)
	return nil
}

func ShowTaskByName(name string) error {
	db, err := db.NewDatabase()
	if err != nil {
		return err
	}
	taskProvider, err := service.NewTaskProvider(db)
	if err != nil {
		return err
	}
	task, err := taskProvider.FindTaskByName(name)
	if err != nil {
		return fmt.Errorf("SERVICE ERROR: %v", err)
	}
	tasks := []repository.Task{task}

	for i := range len(tasks) {
		recordedTime, err := taskProvider.CalculateTaskTime(tasks[i].ID)
		if err != nil {
			return fmt.Errorf("Could not calculate RecordedTime for %d", tasks[i].ID)
		}
		task.RecordedTime = sql.NullInt64{Int64: recordedTime, Valid: true}
	}

	display.PrintTasks(tasks)
	return nil
}

func ShowAllTasks() error {
	db, err := db.NewDatabase()
	if err != nil {
		return err
	}
	taskProvider, err := service.NewTaskProvider(db)
	tasks, err := taskProvider.FindAllTasks()
	if err != nil {
		return fmt.Errorf("SERVICE ERROR: %v", err)
	}

	for i := range len(tasks) {
		recordedTime, err := taskProvider.CalculateTaskTime(tasks[i].ID)
		if err != nil {
			return err
		}
		tasks[i].RecordedTime = sql.NullInt64{Int64: recordedTime, Valid: true}
	}
	display.PrintTasks(tasks)
	return nil
}

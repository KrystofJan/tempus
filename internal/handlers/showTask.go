package handlers

import (
	"database/sql"
	"fmt"

	"github.com/KrystofJan/tempus/internal/display"
	"github.com/KrystofJan/tempus/internal/repository"
	"github.com/KrystofJan/tempus/internal/service"
)

func ShowTaskById(id int64) error {
	taskProvider, err := service.NewTaskProvider()
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
	taskProvider, err := service.NewTaskProvider()
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
	taskProvider, err := service.NewTaskProvider()
	if err != nil {
		return err
	}
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

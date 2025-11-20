package service

import (
	"context"
	"fmt"

	"github.com/KrystofJan/tempus/internal/db"
	"github.com/KrystofJan/tempus/internal/repository"
)

type SwitchService struct {
	ctx  context.Context
	repo *repository.Queries
	db   *db.Database
}

func NewSwitchService(db *db.Database) (*SwitchService, error) {
	ctx := context.Background()
	repo := repository.New(db.Instance)
	return &SwitchService{
		ctx:  ctx,
		repo: repo,
		db:   db,
	}, nil
}

func (self *SwitchService) SwitchTask(taskName string) (*repository.Task, error) {
	tran, err := db.NewTransaction[repository.Task]()
	if err != nil {
		return nil, fmt.Errorf("Error initiating transaction: %v", err)
	}

	var createNewTask bool = false
	task, err := self.repo.FindTaskByName(self.ctx, taskName)
	if err != nil {
		fmt.Printf("Could not find the desired task, creating a new one")
		createNewTask = true
	}
	switchTaskTransaction := func(qtx *repository.Queries) (repository.Task, error) {
		if createNewTask {
			task, err = qtx.AddTask(self.ctx, taskName)
			if err != nil {
				return repository.Task{},
					fmt.Errorf("Error changing task: %v", err)
			}
		}
		current_entry, err := qtx.GetCurrentEntry(self.ctx)
		if err != nil {
			return repository.Task{},
				fmt.Errorf("Cannot query the database: %v", err)
		}
		if current_entry.CurrentEntryID.Valid {
			err = qtx.FinishEntry(self.ctx, current_entry.CurrentEntryID.Int64)
			if err != nil {
				return repository.Task{},
					fmt.Errorf("Error finishing the entry: %v", err)
			}
		}
		_, err = qtx.AddEntry(self.ctx, task.ID)
		if err != nil {
			return repository.Task{},
				fmt.Errorf("Error finishing the entry: %v", err)
		}
		return task, nil
	}

	return tran.PerformTransaction(
		self.repo,
		switchTaskTransaction,
	)
}

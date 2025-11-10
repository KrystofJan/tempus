package service

import (
	"context"
	"fmt"

	"github.com/KrystofJan/tempus/internal/db"
	"github.com/KrystofJan/tempus/internal/repository"
)

type TaskProvider struct {
	ctx  context.Context
	repo *repository.Queries
}

func NewTaskProvider() (*TaskProvider, error) {
	db, err := db.NewDatabase()
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	repo := repository.New(db.Instance)
	return &TaskProvider{
		ctx:  ctx,
		repo: repo,
	}, nil
}

func (self *TaskProvider) CalculateTaskTime(taskId int64) (int64, error) {
	recordedTime, err := self.repo.CalculateTaskTime(self.ctx, taskId)
	if err != nil {
		return 0, fmt.Errorf("Cannot query the database: %v", err)
	}
	return int64(recordedTime.Float64), nil
}

func (self *TaskProvider) FindAllTasks() ([]repository.Task, error) {
	entries, err := self.repo.FindAllTasks(self.ctx)
	if err != nil {
		return nil, fmt.Errorf("Cannot query the database: %v", err)
	}
	return entries, nil
}

func (self *TaskProvider) FindTaskById(id int64) (repository.Task, error) {
	tasks, err := self.repo.FindTaskById(self.ctx, id)
	if err != nil {
		return repository.Task{}, fmt.Errorf("Cannot query the database: %v", err)
	}
	return tasks, nil
}

func (self *TaskProvider) FindTaskByName(name string) (repository.Task, error) {
	db, err := db.NewDatabase()
	if err != nil {
		return repository.Task{}, err
	}
	defer db.Instance.Close()
	tasks, err := self.repo.FindTaskByName(self.ctx, name)
	if err != nil {
		return repository.Task{}, fmt.Errorf("Cannot query the database: %v", err)
	}
	return tasks, nil
}

func (self *TaskProvider) AddTask(name string) (repository.Task, error) {
	task, err := self.repo.AddTask(self.ctx, name)
	if err != nil {
		return repository.Task{}, err
	}
	return task, nil
}

func (self *TaskProvider) ClearTasks() error {
	if err := self.repo.ClearTasks(self.ctx); err != nil {
		return err
	}
	return nil
}

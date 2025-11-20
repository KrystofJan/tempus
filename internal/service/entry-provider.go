package service

import (
	"context"
	"fmt"

	"github.com/KrystofJan/tempus/internal/db"
	"github.com/KrystofJan/tempus/internal/repository"
)

type EntryProvider struct {
	ctx  context.Context
	repo *repository.Queries
}

func NewEntryProvider(db *db.Database) (*EntryProvider, error) {
	ctx := context.Background()
	repo := repository.New(db.Instance)
	return &EntryProvider{
		ctx:  ctx,
		repo: repo,
	}, nil
}

func (self *EntryProvider) FindAllEntries() ([]repository.Entry, error) {
	entries, err := self.repo.FindAllEntries(self.ctx)
	if err != nil {
		return nil, fmt.Errorf("Cannot query the database: %v", err)
	}
	return entries, nil
}

func (self *EntryProvider) FindEntryById(id int64) (repository.Entry, error) {
	entries, err := self.repo.FindEntryById(self.ctx, id)
	if err != nil {
		return repository.Entry{}, fmt.Errorf("Cannot query the database: %v", err)
	}
	return entries, nil
}

func (self *EntryProvider) ClearEntries() error {
	if err := self.repo.ClearEntries(self.ctx); err != nil {
		return err
	}
	return nil
}

func (self *EntryProvider) AddEntry(id int64) (repository.Entry, error) {
	entries, err := self.repo.AddEntry(self.ctx, id)
	if err != nil {
		return repository.Entry{}, fmt.Errorf("Cannot query the database: %v", err)
	}
	return entries, nil
}

func (self *EntryProvider) CalculateEntryTime(id int64) (int64, error) {
	recordedTime, err := self.repo.CalculateEntryTime(self.ctx, id)
	if err != nil {
		return 0, fmt.Errorf("Cannot query the database: %v", err)
	}
	return int64(recordedTime.Float64), nil
}

func (self *EntryProvider) MoveEntry(entryId, taskId int64) error {
	params := repository.MoveEntryParams{TaskID: taskId, ID: entryId}
	return self.repo.MoveEntry(self.ctx, params)
}

func (self *EntryProvider) DeleteEntry(id int64) error {
	return self.repo.DeleteEntry(self.ctx, id)
}

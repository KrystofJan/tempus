package handlers

import (
	"github.com/KrystofJan/tempus/internal/config"
	"github.com/KrystofJan/tempus/internal/errors"
	"github.com/KrystofJan/tempus/internal/repository"
	"github.com/KrystofJan/tempus/internal/service"
)

func EnsureDefaultTaskExists() (*repository.Task, error) {
	cfg, cfgErr := config.Get()
	if cfgErr != nil {
		if cfgErr.ErrorCode != errors.ConfigFileNoExists {
			return nil, cfgErr
		} else {
			if err := GenerateConfig(); err != nil {
				return nil, err
			}
		}
	}

	taskProvider, err := service.NewTaskProvider()
	if err != nil {
		return nil, err
	}

	task, err := taskProvider.FindTaskByName(cfg.DefaultTask)
	if err != nil {
		task, err := taskProvider.AddTask(cfg.DefaultTask)
		if err != nil {
			return nil, err
		}
		return &task, nil
	}

	return &task, nil
}

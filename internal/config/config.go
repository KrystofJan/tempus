package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/KrystofJan/tempus/internal/constants"
	"github.com/KrystofJan/tempus/internal/errors"
	"github.com/KrystofJan/tempus/internal/utils"
)

// NOTE: Get config -> set default task, etc.
type Config struct {
	DefaultTask string `json:"default_task"`
}

func (config *Config) Save() (*Config, error) {
	path, err := utils.GetPath(constants.CONFIG_FOLDER_PATH, constants.CONFIG_FILE_NAME)
	if err != nil {
		return nil, errors.New("Config file not valid", errors.ConfigPath, err)
	}
	marshalledJson, err := json.Marshal(*config)
	if err != nil {
		return nil, err
	}
	err = os.WriteFile(path, marshalledJson, 0644)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func New() (*Config, *errors.ConfigError) {
	path, err := utils.GetPath(constants.CONFIG_FOLDER_PATH, constants.CONFIG_FILE_NAME)
	if err != nil {
		return nil, errors.New("Config file not valid", errors.ConfigPath, err)
	}
	_, error := os.Stat(path)
	if !os.IsNotExist(error) {
		fmt.Printf("Config file does not exist, creating a new one")
	}
	config := &Config{
		DefaultTask: constants.GENERATED_DEFAULT_TASK_NAME,
	}
	config, err = config.Save()
	if err != nil {
		return nil, errors.New("Saving to the config file failed", errors.ErrorSavingToConfigFile, err)
	}
	return config, nil
}

func Get() (*Config, *errors.ConfigError) {
	path, err := utils.GetPath(constants.CONFIG_FOLDER_PATH, constants.CONFIG_FILE_NAME)
	if err != nil {
		return nil, errors.New("Config file not valid", errors.ConfigPath, err)
	}
	_, error := os.Stat(path)

	if os.IsNotExist(error) {
		return nil, errors.New("Config file does not exists", errors.ConfigFileExists, error)
	}

	config, configErr := readConfigFromFile(path)
	if configErr != nil {
		return nil, errors.New("Reading from the config file failed", errors.ErrorReadingConfigFile, err)
	}
	return config, nil
}

func readConfigFromFile(path string) (*Config, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var config Config

	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func Delete() error {
	path, err := utils.GetPath(constants.CONFIG_FOLDER_PATH, constants.CONFIG_FILE_NAME)
	if err != nil {
		return errors.New("Config file not valid", errors.ConfigPath, err)
	}
	_, error := os.Stat(path)

	if os.IsNotExist(error) {
		fmt.Print("Config file does not exists, won't delete\n")
		return nil
	}
	if err = os.Remove(path); err != nil {
		return err
	}
	return nil
}

func (cfg *Config) ToString() string {
	out, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(out)
}

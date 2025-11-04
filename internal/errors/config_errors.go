package errors

import (
	"fmt"
)

type ConfigErrorCode int

const (
	ConfigFileNoExists ConfigErrorCode = iota
	ConfigFileExists
	ErrorSavingToConfigFile
	ErrorReadingConfigFile
	NewConfigConflict
	ConfigPath
)

type ConfigError struct {
	s                string
	ErrorCode        ConfigErrorCode
	originatingError error
}

func New(message string, errorCode ConfigErrorCode, err error) *ConfigError {
	return &ConfigError{
		s:                message,
		ErrorCode:        errorCode,
		originatingError: err,
	}
}

func (cfg ConfigError) Error() string {
	return fmt.Sprintf("CONFIG ERROR: %s\nOriginalError:%v\n", cfg.s, cfg.originatingError)
}

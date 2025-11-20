package constants

type EnvironmentType string

var (
	Production  EnvironmentType = "production"
	Development EnvironmentType = "development"
	Testing     EnvironmentType = "testing"
)

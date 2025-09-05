package application

type EnvironmentType string

const (
	EnvironmentDevelopment = "development"
)

type Config struct {
	Env      EnvironmentType `mapstructure:"env"`
	Database DatabaseConfig  `mapstructure:"database"`
}

type DatabaseConfig struct {
	URL string `mapstructure:"db_url"`
}

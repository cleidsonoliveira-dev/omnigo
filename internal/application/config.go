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
	// ex: "postgres://user:pass@localhost:5432/dbname"
	URL                      string `mapstructure:"db_url"`
	MaxConn                  int32  `mapstructure:"max_conn"`
	MinConn                  int32  `mapstructure:"min_conn"`
	MaxConnIdleTimeMinutes   int32  `mapstructure:"max_conn_idle_time_minutes"`
	MaxConnLifetimeMinutes   int32  `mapstructure:"max_conn_lifetime_minutes"`
	HealthCheckPeriodMinutes int32  `mapstructure:"health_check_period_minutes"`
}

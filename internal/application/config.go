// Copyright 2025 The Omnigo Authors. All Rights Reserved.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
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

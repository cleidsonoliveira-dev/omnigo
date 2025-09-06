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
package infra

import (
	"fmt"
	"strings"

	"github.com/cleidsonoliveira-dev/omnigo/internal/application"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func LoadConfig(path string, logger *zap.Logger) (application.Config, error) {
	var cfg application.Config

	if path == "" {
		path = "."
	}

	viper.SetConfigName("omnigo")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)

	viper.SetEnvPrefix("omnigo")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// Default Values
	viper.SetDefault("env", application.EnvironmentDevelopment)
	viper.SetDefault("database.max_conn", 5)
	viper.SetDefault("database.min_conn", 2)
	viper.SetDefault("database.max_conn_idle_time_minutes", 5)
	viper.SetDefault("database.max_conn_lifetime_minutes", 30)
	viper.SetDefault("database.health_check_period_minutes", 1)

	if err := viper.ReadInConfig(); err != nil {
		logger.Info("Application configuration not found", zap.Error(err))
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return cfg, fmt.Errorf("error during application config decode: %w", err)
	}

	logger.Info(fmt.Sprintf("Current Environment %s", cfg.Env))

	logger.Info("Configuration loaded successfully.")
	return cfg, nil
}

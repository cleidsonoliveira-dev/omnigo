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

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
	"context"
	"time"

	"github.com/cleidsonoliveira-dev/omnigo/internal/application"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

func makePgx(cfg application.Config, logger *zap.Logger) (*pgxpool.Pool, error) {
	dsn := cfg.Database.URL

	pgxCfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	pgxCfg.MaxConns = cfg.Database.MaxConn
	pgxCfg.MinConns = cfg.Database.MinConn
	pgxCfg.MaxConnLifetime = time.Duration(cfg.Database.MaxConnLifetimeMinutes) * time.Minute
	pgxCfg.MaxConnIdleTime = time.Duration(cfg.Database.MaxConnIdleTimeMinutes) * time.Minute
	pgxCfg.HealthCheckPeriod = time.Duration(cfg.Database.HealthCheckPeriodMinutes) * time.Minute

	pool, err := pgxpool.NewWithConfig(context.Background(), pgxCfg)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}

	logger.Info("Postgres pool started successfully")
	return pool, nil
}

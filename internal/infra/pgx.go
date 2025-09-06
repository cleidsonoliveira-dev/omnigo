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

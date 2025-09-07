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
	"embed"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.SQL
var migrationsFS embed.FS

func MigrateUp(ctx context.Context, pool *pgxpool.Pool) error {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("acquiring pgx connection: %w", err)
	}
	// Copy of ConnConfig for Goose
	connCfg := *conn.Conn().Config()
	conn.Release()

	sqlDB := stdlib.OpenDB(connCfg)
	defer sqlDB.Close()

	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("setting goose dialect: %w", err)
	}
	goose.SetBaseFS(migrationsFS)

	const migrationsDir = "migrations"

	if err := goose.UpContext(ctx, sqlDB, migrationsDir); err != nil {
		return fmt.Errorf("running goose up: %w", err)
	}
	return nil
}

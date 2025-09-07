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
    "fmt"
    "time"

    sq "github.com/Masterminds/squirrel"
    "github.com/cleidsonoliveira-dev/omnigo/internal/core"
    "github.com/google/uuid"
    "github.com/jackc/pgx/v5/pgxpool"
)

type PgxChannelRepository struct {
	pool *pgxpool.Pool
}

var (
    repoChannelTable = "core.channels"
)

func (repo *PgxChannelRepository) Create(ctx context.Context, channel *core.Channel) (*core.Channel, error) {
    now := time.Now()

    sqlBuilder := sq.Insert(repoChannelTable).
        Columns("provider_id", "name", "description", "created_at", "updated_at").
        Values(channel.ProviderID, channel.Name, channel.Description, now, now).
        Suffix("RETURNING id, created_at, updated_at").
        PlaceholderFormat(sq.Dollar)

    query, args, err := sqlBuilder.ToSql()
    if err != nil {
        return nil, fmt.Errorf("error building SQL: %w", err)
    }

    var id uuid.UUID
    var createdAt, updatedAt time.Time
    if err := repo.pool.QueryRow(ctx, query, args...).Scan(&id, &createdAt, &updatedAt); err != nil {
        return nil, fmt.Errorf("error executing insert: %w", err)
    }

    channel.ID = id
    channel.CreatedAt = createdAt
    channel.UpdatedAt = updatedAt
    return channel, nil
}

func (repo *PgxChannelRepository) GetByID(ctx context.Context, id uuid.UUID) (*core.Channel, error) {
    sqlBuilder := sq.Select(
        "id",
        "provider_id",
        "name",
        "COALESCE(description, '') AS description",
        "created_at",
        "updated_at",
    ).From(repoChannelTable).
        Where(sq.Eq{"id": id}).
        PlaceholderFormat(sq.Dollar)

    query, args, err := sqlBuilder.ToSql()
    if err != nil {
        return nil, fmt.Errorf("error building SQL: %w", err)
    }

    var ch core.Channel
    if err := repo.pool.QueryRow(ctx, query, args...).Scan(
        &ch.ID,
        &ch.ProviderID,
        &ch.Name,
        &ch.Description,
        &ch.CreatedAt,
        &ch.UpdatedAt,
    ); err != nil {
        return nil, fmt.Errorf("error querying channel by id: %w", err)
    }
    return &ch, nil
}

func (repo *PgxChannelRepository) Save(ctx context.Context, channel *core.Channel) (*core.Channel, error) {
    now := time.Now()

    sqlBuilder := sq.Update(repoChannelTable).
        Set("provider_id", channel.ProviderID).
        Set("name", channel.Name).
        Set("description", channel.Description).
        Set("updated_at", now).
        Where(sq.Eq{"id": channel.ID}).
        Suffix("RETURNING created_at, updated_at").
        PlaceholderFormat(sq.Dollar)

    query, args, err := sqlBuilder.ToSql()
    if err != nil {
        return nil, fmt.Errorf("error building SQL: %w", err)
    }

    var createdAt, updatedAt time.Time
    if err := repo.pool.QueryRow(ctx, query, args...).Scan(&createdAt, &updatedAt); err != nil {
        return nil, fmt.Errorf("error executing update: %w", err)
    }

    channel.CreatedAt = createdAt
    channel.UpdatedAt = updatedAt
    return channel, nil
}

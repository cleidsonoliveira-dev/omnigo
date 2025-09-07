//go:build repo_integration
// +build repo_integration

package infra

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/cleidsonoliveira-dev/omnigo/internal/application"
	"github.com/cleidsonoliveira-dev/omnigo/internal/core"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

var (
	testLogger *zap.Logger
	testCfg    application.Config
	testPool   *pgxpool.Pool
)

func TestMain(m *testing.M) {
	var err error
	testLogger, err = zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	defer testLogger.Sync()

	testCfg, err = LoadConfig("omnigo_test", "", testLogger)
	if err != nil {
		testLogger.Fatal("failed to load config", zap.Error(err))
	}

	testPool, err = makePgx(testCfg, testLogger)
	if err != nil {
		testLogger.Fatal("failed to init pgx pool", zap.Error(err))
	}

	// Run migrations once for the whole suite
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := MigrateUp(ctx, testPool); err != nil {
		testLogger.Fatal("failed to run migrations", zap.Error(err))
	}

	code := m.Run()

	// Teardown: reset DB after all tests
	ctxDown, cancelDown := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancelDown()
	if err := MigrateDown(ctxDown, testPool); err != nil {
		testLogger.Error("failed to run migrations down", zap.Error(err))
	}

	// Close pool explicitly before exiting
	testPool.Close()

	os.Exit(code)
}

func TestPgxChannelRepository_CRUD(t *testing.T) {
	ctx := context.Background()
	repo := &PgxChannelRepository{pool: testPool}

	// Create
	ch := &core.Channel{
		ProviderID:  "prov-" + uuid.NewString(),
		Name:        "Test Channel",
		Description: "A channel used in integration tests",
	}
	created, err := repo.Create(ctx, ch)
	if err != nil {
		t.Fatalf("Create() error: %v", err)
	}
	if created.ID == uuid.Nil {
		t.Fatalf("Create() expected non-nil ID")
	}
	if created.CreatedAt.IsZero() || created.UpdatedAt.IsZero() {
		t.Fatalf("Create() expected timestamps to be set")
	}

	// GetByID
	got, err := repo.GetByID(ctx, created.ID)
	if err != nil {
		t.Fatalf("GetByID() error: %v", err)
	}
	if got.ID != created.ID || got.ProviderID != ch.ProviderID || got.Name != ch.Name || got.Description != ch.Description {
		t.Fatalf("GetByID() returned unexpected record: %+v", got)
	}

	// Save (update)
	prevUpdatedAt := created.UpdatedAt
	created.Name = "Updated Name"
	created.Description = "Updated Description"

	saved, err := repo.Save(ctx, created)
	if err != nil {
		t.Fatalf("Save() error: %v", err)
	}
	if !saved.UpdatedAt.After(prevUpdatedAt) && !saved.UpdatedAt.Equal(prevUpdatedAt) {
		t.Fatalf("Save() expected UpdatedAt to be >= previous UpdatedAt")
	}

	got2, err := repo.GetByID(ctx, created.ID)
	if err != nil {
		t.Fatalf("GetByID() after update error: %v", err)
	}
	if got2.Name != "Updated Name" || got2.Description != "Updated Description" {
		t.Fatalf("GetByID() after update mismatch, got: %+v", got2)
	}
}

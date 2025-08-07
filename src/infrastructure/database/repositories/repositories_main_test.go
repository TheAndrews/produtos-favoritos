package repositories

import (
	"context"
	"fmt"
	"log"
	"os"
	"produtos-favoritos/src/infrastructure/database/migrations"
	"testing"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Global DB instance for tests
var TestDB *gorm.DB
var postgresContainer testcontainers.Container

func TestMain(m *testing.M) {
	ctx := context.Background()

	// Start PostgreSQL container
	req := testcontainers.ContainerRequest{
		Image:        "postgres:15",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_USER":     "testuser",
			"POSTGRES_PASSWORD": "testpass",
			"POSTGRES_DB":       "testdb",
		},
		WaitingFor: wait.ForListeningPort("5432/tcp"),
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		log.Fatalf("Could not start container: %s", err)
	}
	postgresContainer = container

	// Get host and port
	host, _ := container.Host(ctx)
	port, _ := container.MappedPort(ctx, "5432")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=testuser dbname=testdb password=testpass sslmode=disable TimeZone=UTC",
		host, port.Port(),
	)

	// Connect to DB using GORM
	TestDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("Failed to connect to container DB: %s", err)
	}

	// Run migrations
	migrations.RunMigrations(TestDB)

	// Run tests
	code := m.Run()

	// Cleanup
	if err := container.Terminate(ctx); err != nil {
		log.Printf("Failed to stop container: %s", err)
	}

	os.Exit(code)
}

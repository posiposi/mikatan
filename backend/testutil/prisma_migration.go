package testutil

import (
	"fmt"
	"os"
	"os/exec"
)

func RunPrismaMigrationForTest() error {
	schemaPath := "../infrastructure/prisma/schema.prisma"

	if err := resetTestDatabase(); err != nil {
		return fmt.Errorf("failed to reset test database: %w", err)
	}

	cmd := exec.Command("go", "run", "github.com/steebchen/prisma-client-go", "migrate", "deploy", "--schema", schemaPath)
	cmd.Env = append(os.Environ(),
		"DATABASE_URL="+getTestDatabaseURL(),
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to run prisma migration: %w, output: %s", err, string(output))
	}

	return nil
}

func resetTestDatabase() error {
	cmd := exec.Command("go", "run", "github.com/steebchen/prisma-client-go", "migrate", "reset", "--force", "--schema", "../infrastructure/prisma/schema.prisma")
	cmd.Env = append(os.Environ(),
		"DATABASE_URL="+getTestDatabaseURL(),
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to reset prisma database: %w, output: %s", err, string(output))
	}

	return nil
}

func getTestDatabaseURL() string {
	mysqlPassword := os.Getenv("MYSQL_TEST_ROOT_PASSWORD")
	mysqlHost := os.Getenv("MYSQL_TEST_HOST")
	mysqlDatabase := os.Getenv("MYSQL_TEST_DATABASE")
	return fmt.Sprintf("mysql://root:%s@%s:3306/%s", mysqlPassword, mysqlHost, mysqlDatabase)
}

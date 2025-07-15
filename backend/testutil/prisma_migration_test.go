package testutil

import (
	"os"
	"testing"
)

func TestGetTestDatabaseURL(t *testing.T) {
	// 元の環境変数を保存
	originalPassword := os.Getenv("MYSQL_TEST_ROOT_PASSWORD")
	originalHost := os.Getenv("MYSQL_TEST_HOST")
	originalDatabase := os.Getenv("MYSQL_TEST_DATABASE")

	// テスト後に元の環境変数を復元
	defer func() {
		os.Setenv("MYSQL_TEST_ROOT_PASSWORD", originalPassword)
		os.Setenv("MYSQL_TEST_HOST", originalHost)
		os.Setenv("MYSQL_TEST_DATABASE", originalDatabase)
	}()

	t.Run("正常なケース", func(t *testing.T) {
		os.Setenv("MYSQL_TEST_ROOT_PASSWORD", "test_pass")
		os.Setenv("MYSQL_TEST_HOST", "test_host")
		os.Setenv("MYSQL_TEST_DATABASE", "test_db")

		url, err := getTestDatabaseURL()
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}

		expected := "mysql://root:test_pass@test_host:3306/test_db"
		if url != expected {
			t.Errorf("Expected %s, got %s", expected, url)
		}
	})

	t.Run("環境変数が一つ不足", func(t *testing.T) {
		os.Setenv("MYSQL_TEST_ROOT_PASSWORD", "test_pass")
		os.Setenv("MYSQL_TEST_HOST", "test_host")
		os.Unsetenv("MYSQL_TEST_DATABASE")

		_, err := getTestDatabaseURL()
		if err == nil {
			t.Error("Expected error for missing environment variable")
		}

		expectedMsg := "missing required environment variables: [MYSQL_TEST_DATABASE]"
		if err.Error() != expectedMsg {
			t.Errorf("Expected error message '%s', got '%s'", expectedMsg, err.Error())
		}
	})

	t.Run("複数の環境変数が不足", func(t *testing.T) {
		os.Unsetenv("MYSQL_TEST_ROOT_PASSWORD")
		os.Unsetenv("MYSQL_TEST_HOST")
		os.Setenv("MYSQL_TEST_DATABASE", "test_db")

		_, err := getTestDatabaseURL()
		if err == nil {
			t.Error("Expected error for missing environment variables")
		}

		expectedMsg := "missing required environment variables: [MYSQL_TEST_ROOT_PASSWORD MYSQL_TEST_HOST]"
		if err.Error() != expectedMsg {
			t.Errorf("Expected error message '%s', got '%s'", expectedMsg, err.Error())
		}
	})

	t.Run("全ての環境変数が不足", func(t *testing.T) {
		os.Unsetenv("MYSQL_TEST_ROOT_PASSWORD")
		os.Unsetenv("MYSQL_TEST_HOST")
		os.Unsetenv("MYSQL_TEST_DATABASE")

		_, err := getTestDatabaseURL()
		if err == nil {
			t.Error("Expected error for missing environment variables")
		}

		expectedMsg := "missing required environment variables: [MYSQL_TEST_ROOT_PASSWORD MYSQL_TEST_HOST MYSQL_TEST_DATABASE]"
		if err.Error() != expectedMsg {
			t.Errorf("Expected error message '%s', got '%s'", expectedMsg, err.Error())
		}
	})
}
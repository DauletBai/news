package database

import (
    "testing"
    "database/sql"
    _ "github.com/lib/pq"
)

func TestConnect(t *testing.T) {
    // Подключение к тестовой базе данных (например, с другой конфигурацией)
    connStr := "host=localhost port=5432 user=test_user password=test_pass dbname=test_db sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        t.Fatalf("Failed to connect to the test database: %v", err)
    }
    defer db.Close()

    // Проверка подключения
    if err = db.Ping(); err != nil {
        t.Fatalf("Ping to test database failed: %v", err)
    }
}
package database

import (
    "database/sql"
    "fmt"
    "log"
    _ "github.com/lib/pq"
    "news/internal/config"
)

func Connect(cfg config.DBConfig) *sql.DB {
    dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)

    db, err := sql.Open("postgres", dsn)
    if err != nil {
        log.Fatalf("cannot connect to db: %v", err)
    }

    if err = db.Ping(); err != nil {
        log.Fatalf("cannot ping db: %v", err)
    }

    log.Println("Connected to the database!")
    return db
}
package config

import (
    "github.com/joho/godotenv"
    "os"
)

type Config struct {
    ServerPort string
    DB         DBConfig
}

type DBConfig struct {
    Host     string
    Port     string
    User     string
    Password string
    Name     string
}

func LoadConfig(path string) (*Config, error) {
    err := godotenv.Load(path)
    if err != nil {
        return nil, err
    }

    cfg := &Config{
        ServerPort: os.Getenv("SERVER_PORT"),
        DB: DBConfig{
            Host:     os.Getenv("DB_HOST"),
            Port:     os.Getenv("DB_PORT"),
            User:     os.Getenv("DB_USER"),
            Password: os.Getenv("DB_PASSWORD"),
            Name:     os.Getenv("DB_NAME"),
        },
    }

    return cfg, nil
}
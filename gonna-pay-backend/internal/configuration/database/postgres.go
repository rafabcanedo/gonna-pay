package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func InitPostgres() (*sql.DB, error) {
    host := getEnv("DB_HOST", "localhost")
    port := getEnv("DB_PORT", "5432")
    user := getEnv("DB_USER", "postgres")
    password := getEnv("DB_PASSWORD", "postgres")
    dbname := getEnv("DB_NAME", "internal")
    sslmode := getEnv("DB_SSLMODE", "disable")

    dsn := fmt.Sprintf(
        "postgres://%s:%s@%s:%s/%s?sslmode=%s",
        user, password, host, port, dbname, sslmode,
    )

    db, err := sql.Open("pgx", dsn)
    if err != nil {
        return nil, fmt.Errorf("error opening connection to postgres: %w", err)
    }

    db.SetMaxOpenConns(25)
    db.SetMaxIdleConns(25)
    db.SetConnMaxLifetime(5 * time.Minute)

    if err := retryPing(db, 10, 2*time.Second); err != nil {
        return nil, err
    }

    DB = db
    log.Println("Connected to Postgres successfully 📊")
    return db, nil
}

func retryPing(db *sql.DB, attempts int, delay time.Duration) error {
    var err error
    for i := 0; i < attempts; i++ {
        err = db.Ping()
        if err == nil {
            return nil
        }
        log.Printf("Postgres not ready yet, attempt %d/%d - error: %v", i+1, attempts, err)
        time.Sleep(delay)
    }
    return fmt.Errorf("could not connect to Postgres after %d attempts: %w", attempts, err)
}

func getEnv(key, defaultValue string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return defaultValue
}

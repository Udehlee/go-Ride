package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func ConnectDB(config Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable",
		config.Username, config.Password, config.Host, config.Port, config.DbName)

	Conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to create database connection: %w", err)
	}

	if err := runMigrations(Conn); err != nil {
		log.Fatalf("Migration unsuccessful: %v", err)
	}

	if err := Conn.Ping(); err != nil {
		return nil, fmt.Errorf("connection not alive: %w", err)
	}

	fmt.Println("connected successfully")
	return Conn, nil
}

func runMigrations(db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("failed to create database driver instance: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://./db/migrations",
		"postgres", driver)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to run up migrations: %w", err)
	}
	log.Println("Migrations applied successfully!")
	return nil
}

func InitDB(config Config) (PgConn, error) {
	pg := PgConn{}

	db, err := ConnectDB(config)
	if err != nil {
		log.Fatal(err)
	}

	pg.Conn = db
	return pg, nil
}

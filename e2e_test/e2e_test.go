package e2e_test

import (
	"database/sql"
	"embed"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	log "github.com/sirupsen/logrus"
	"testing"
)

var db *sql.DB

//go:embed migrations/*.sql
var embedMigrations embed.FS

func TestMain(m *testing.M) {
	goose.SetBaseFS(embedMigrations)

	err := goose.SetDialect("postgres")
	if err != nil {
		panic(err)
	}
	db, err = sql.Open("postgres", "postgres://postgres:7dgvJVDJvh254aqOpfd@postgres:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatalf("Could not connected: %s", err)
	}

	err = goose.UpTo(db, "migrations", 20231206192143)
	if err != nil {
		log.Warnf("Error migration: %s", err)
	}
	m.Run()

	err = goose.DownTo(db, "migrations", 20231206192143)
	if err != nil {
		log.Warnf("Error migration: %s", err)
	}
}

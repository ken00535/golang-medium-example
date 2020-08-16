package migrate

import (
	"log"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
)

type Migration struct {
	client *migrate.Migrate
}

func New() *Migration {
	m := Migration{}
	path, err := os.Executable()
	if err != nil {
		log.Panic(err)
	}
	path = "file://" + filepath.Join(path, "..", "..", "migrations")
	m.client, err = migrate.New(path, "postgres://postgres:@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Panic(err)
	}
	return &m
}

// Up to newest version
func (m *Migration) Up() {
	if err := m.client.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}

// Down to oldest current
func (m *Migration) Down() {
	if err := m.client.Down(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}

// Force sets a migeration version to
func (m *Migration) Force(version int) {
	if err := m.client.Force(version); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}

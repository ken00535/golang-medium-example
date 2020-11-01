package main

import (
	"example/pkg/migrate"
	"flag"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {

	var version int
	var up, down bool
	flag.IntVar(&version, "force", -1, "sets migration as specific version")
	flag.BoolVar(&up, "up", false, "up to newest")
	flag.BoolVar(&down, "down", false, "down to oldest")
	flag.Parse()

	client := migrate.New()
	if version != -1 {
		client.Force(version)
	}
	if up {
		client.Up()
	}
	if down {
		client.Down()
	}
}

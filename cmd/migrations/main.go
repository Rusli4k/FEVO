package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/rusli4k/fevo/app/repository/pg"
	"github.com/rusli4k/fevo/cfg"

	migrate "github.com/rubenv/sql-migrate"
)

const (
	up           = "up"
	down         = "down"
	migrationDir = "./migration"
)

func main() {
	config, err := cfg.GetConfig()
	if err != nil {
		log.Printf("Failed to get config: %+v", err)

		return
	}

	db, err := pg.ConnectDB(config.DB)
	if err != nil {
		log.Printf("%+v", err)

		return
	}

	direction := flag.String("migrate", "", "applying migration direction")
	flag.Parse()

	if *direction != up && *direction != down {
		log.Printf("Wrong flag provided, choose '-migrate %s' or '-migrate %s'\n", up, down)

		return
	}

	if err := migrateDB(db, *direction); err != nil {
		log.Printf("Failed making migrations: %v", err)
	}
}

// MigrateDB executes migrations.
func migrateDB(db *sql.DB, direction string) error {
	migrations := &migrate.FileMigrationSource{
		Dir: migrationDir,
	}
	var dir migrate.MigrationDirection
	if direction == down {
		dir = 1
	}

	log.Printf("Starting applying migrations '%s'...", direction)

	n, err := migrate.Exec(db, "postgres", migrations, dir)
	if err != nil {
		return fmt.Errorf("migration up failed: %w", err)
	}

	log.Printf("The number of applied migration is: %d", n)

	return nil
}

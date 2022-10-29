package main

import (
	"fmt"
	"log"

	"github.com/rusli4k/fevo/app/repository/pg"
	"github.com/rusli4k/fevo/app/transport/rest"
	"github.com/rusli4k/fevo/app/usecase"
	"github.com/rusli4k/fevo/cfg"
)

func main() {
	if err := Run(); err != nil {
		log.Fatalf("Cann't run server: %v", err)
	}
}

// Run will bind our layers all together.
func Run() error {
	config, err := cfg.GetConfig()
	if err != nil {
		return fmt.Errorf("failed to get config: %w", err)
	}

	db, err := pg.ConnectDB(config.DB)
	if err != nil {
		return fmt.Errorf("error connecting to database on host: %s, port: %s, with error: %w", config.DB.Host, config.DB.Port, err)
	}

	repo := pg.NewRepo(db)

	handlers := rest.Handlers{
		TAHandler: rest.NewTAHandler(usecase.NewTA(repo)),
	}

	srv := rest.NewServer(config, handlers)

	if err := srv.Run(); err != nil {
		return fmt.Errorf("error loading server: %w", err)
	}

	return nil
}

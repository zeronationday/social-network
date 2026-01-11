package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/zeronationday/social-network/internal/env"
)

func main() {
	ctx := context.Background()

	err := godotenv.Load()
	if err != nil {
		log.Panic(err)
	}

	cfg := config{
		addr: env.GetString("ADDR"),
		db: dbConfig{
			dsn: env.GetString("GOOSE_DBSTRING"),
		},
	}

	conn, err := pgx.Connect(ctx, cfg.db.dsn)
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close(ctx)

	app := &application{
		config: cfg,
		db:     conn,
	}

	h := app.mount()
	log.Fatal(app.run(h))
}

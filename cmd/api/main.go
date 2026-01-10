package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/zeronationday/social-network/internal/db"
	"github.com/zeronationday/social-network/internal/env"
	"github.com/zeronationday/social-network/internal/storage"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Panic(err)
	}

	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://user:password@localhost/social_network?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetInt("DB_MAX_IDLE_TIME", 15),
		},
	}

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	log.Println("database connected successfully")

	storage := storage.NewStorage(db)

	app := &application{
		config:  cfg,
		storage: storage,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}

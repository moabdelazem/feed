package main

import (
	"log"

	"github.com/moabdelazem/feed/internal/db"
	"github.com/moabdelazem/feed/internal/env"
	"github.com/moabdelazem/feed/internal/store"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8000"),
		dbConfig: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/social?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	db, err := db.New(cfg.dbConfig.addr, cfg.dbConfig.maxIdleTime, cfg.dbConfig.maxOpenConns, cfg.dbConfig.maxIdleConns)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()
	log.Println("Database connection is established.")

	store := store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}

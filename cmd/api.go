package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
	repo "github.com/zeronationday/social-network/internal/adapters/postgresql/sqlc"
	"github.com/zeronationday/social-network/internal/posts"
	"github.com/zeronationday/social-network/internal/users"
)

type application struct {
	config config
	db     *pgx.Conn
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	dsn string
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(time.Minute))

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	userService := users.NewService(repo.New(app.db))
	userHandler := users.NewHandler(userService)
	r.Get("/users", userHandler.ListUsers)
	r.Get("/users/{id}", userHandler.FindUserByID)
	r.Post("/users", userHandler.CreateUser)
	r.Put("/users/{id}", userHandler.UpdateUser)

	postService := posts.NewService(repo.New(app.db))
	postHandler := posts.NewHandler(postService)
	r.Get("/posts/user/{user_id}", postHandler.ListPostsByUserID)

	return r
}

func (app *application) run(h http.Handler) error {
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      h,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("server has started at %s", app.config.addr)

	return srv.ListenAndServe()
}

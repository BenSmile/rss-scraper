package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/bensmile/rssaggregator/internal/config"
	"github.com/bensmile/rssaggregator/internal/database"
	"github.com/bensmile/rssaggregator/internal/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	godotenv.Load(".env")
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not in the env")
	}
	fmt.Printf("Server is running on port %s\n", portString)

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found in the env")
	}

	conn, err := sql.Open("postgres", fmt.Sprintf("%s?sslmode=disable", dbURL))
	if err != nil {
		log.Fatalf("Failed to connect to the database : %v", err)
	}
	if err := conn.Ping(); err != nil {
		log.Fatal("Failed to ping the db")
	}

	queries := database.New(conn)

	apiCfg := &config.ApiConfig{
		DB: queries,
	}

	go startScraping(
		queries,
		10,
		time.Minute,
	)

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()

	v1Router.Get("/health", handlers.HandlerReadiness)
	v1Router.Get("/error", handlers.HandlerErr)

	handlers := handlers.NewHandler(apiCfg)
	{
		v1Router.Post("/users", handlers.HandlerCreateUser)
		v1Router.Get("/users", handlers.MiddlewareAuth(handlers.HandlerGetUser))
		v1Router.Get("/users/posts", handlers.MiddlewareAuth(handlers.HandlerGetPostsByUser))
	}
	{
		v1Router.Post("/feeds", handlers.MiddlewareAuth(handlers.HandlerCreateFeed))
		v1Router.Get("/feeds", handlers.HandlerGetFeeds)
	}
	{
		v1Router.Post("/feed-follows", handlers.MiddlewareAuth(handlers.HandlerCreateFeedFollow))
		v1Router.Get("/feed-follows", handlers.MiddlewareAuth(handlers.HandlerGetFeedsFollows))
		v1Router.Delete("/feed-follows/{feed_follow_id}", handlers.MiddlewareAuth(handlers.HandlerUnfollowFeed))
	}

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

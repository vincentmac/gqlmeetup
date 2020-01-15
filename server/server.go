package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-pg/pg/v9"
	"github.com/joho/godotenv"
	"github.com/rs/cors"

	"github.com/vincentmac/gqlmeetup/graphql"
	customMiddleware "github.com/vincentmac/gqlmeetup/middleware"
	"github.com/vincentmac/gqlmeetup/postgres"
)

const defaultPort = "8080"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DB := postgres.New(&pg.Options{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_DATABASE"),
	})

	defer DB.Close()

	DB.AddQueryHook(postgres.DBLogger{})

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	userRepo := postgres.UsersRepo{DB: DB}

	router := chi.NewRouter()

	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{fmt.Sprintf("http://%v:%v", "localhost", port)},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(customMiddleware.AuthMiddleware(userRepo))

	c := graphql.Config{Resolvers: &graphql.Resolver{
		MeetupsRepo: postgres.MeetupsRepo{DB: DB},
		UsersRepo:   userRepo,
	}}

	queryHandler := handler.GraphQL(graphql.NewExecutableSchema(c))

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", graphql.DataloaderMiddleware(DB, queryHandler)) // add Dataloader Middleware

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

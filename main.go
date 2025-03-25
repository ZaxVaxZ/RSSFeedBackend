package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ZaxVaxZ/RSSFeedBackend/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load(".env")

	address := os.Getenv("ADDRESS")
	port := os.Getenv("PORT")	
	if port == "" {
		log.Fatal("PORT is not found in the environment")
	}
	
	dbURL := os.Getenv("DB_URL")
	
	if dbURL == "" {
		log.Fatal("Database URL is not found in the environment")
	}
	
	conn, err := sql.Open("postgres", dbURL)
	
	if err != nil {
		log.Fatal("Can't connect to Database")
	}
	
	apiCfg := apiConfig {
		DB: database.New(conn),
	}

	switch address {
		case "":
			fmt.Printf("Server listening on port %v...\n", port)
		default:
			fmt.Printf("Server listening on %v:%v...\n", address, port)
	}
	
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)
	v1Router.Post("/users", apiCfg.handlerCreateUser)
	v1Router.Get("/users", apiCfg.handlerGetUser)
	v1Router.Delete("/users", apiCfg.handlerDeleteUser)

	router.Mount("/v1", v1Router)

	srv := &http.Server {
		Handler: router,
		Addr: address + ":" + port,
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
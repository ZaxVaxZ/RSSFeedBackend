package main

import (
	"fmt"
	"log"
	"os"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func main() {
	fmt.Println("Test")

	godotenv.Load(".env")

	address := os.Getenv("ADDRESS")
	port := os.Getenv("PORT")
	
	if port == "" {
		log.Fatal("PORT is not found in the environment")
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

	router.Mount("/v1", v1Router)

	srv := &http.Server {
		Handler: router,
		Addr: address + ":" + port,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
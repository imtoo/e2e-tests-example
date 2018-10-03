package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/imtoo/e2e-tests-example/addarticle"
	"github.com/imtoo/e2e-tests-example/config"
	"github.com/imtoo/e2e-tests-example/database"
	"github.com/imtoo/e2e-tests-example/models"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"

	_ "github.com/lib/pq"
)

func main() {
	database.SetupDB()
	db := database.OpenDB()
	store := &models.StoreType{DB: db}

	router := mux.NewRouter()
	pathAndPort := config.EnvPathAndPort

	// ARTICLES
	router.HandleFunc(
		config.RouteArticleAdd,
		addarticle.StoreType{Store: store}.Handler,
	).Methods("POST")

	// MIDDLEWARES
	// CORS
	corsOptions := cors.New(cors.Options{
		AllowedHeaders:   []string{"*"},
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})
	routerWithCORS := corsOptions.Handler(router)
	// Logging
	routerWithLogging := handlers.LoggingHandler(os.Stdout, routerWithCORS)

	fmt.Printf("Listening at http://localhost%s/\n", config.EnvPort)

	// Run server
	log.Fatal(http.ListenAndServe(pathAndPort, routerWithLogging))
}

package main

import (
	"log"
	"net/http"
	"os"

	"S3_FriendManagement_Graphql/config"
	"S3_FriendManagement_Graphql/routes"
	"github.com/joho/godotenv"
)

const defaultPort = "8080"

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	//Database
	confDB := config.Database{}
	database := confDB.InitDatabase()

	// Routes
	router := routes.Routes{
		Db: database,
	}
	server := router.Register()

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, server))
}

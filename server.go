package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")

	// Initialize router
	r := mux.NewRouter()

	fmt.Printf("Server listening on port %s", PORT)
	http.ListenAndServe(fmt.Sprintf(":%s", PORT), r)
}

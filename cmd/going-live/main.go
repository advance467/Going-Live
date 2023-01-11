package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

var (
	port = "3001"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, MainRoutes())
}

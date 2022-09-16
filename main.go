package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	// port env start
	port := ""
	port = os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	// port env end

	MainRouter := RouterReady()
	err := http.ListenAndServe(":"+port, MainRouter)
	if err != nil {
		panic(err)
	}
	fmt.Println("Server has been started")
}

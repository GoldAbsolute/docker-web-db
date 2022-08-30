package main

import (
	"fmt"
	"net/http"
	"os"
)

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

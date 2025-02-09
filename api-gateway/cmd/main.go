package main

import (
	"api-gateway/api"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := api.NewRouter()

	fmt.Println("REST server running...")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("failed to start REST server: %v", err)
	}
}

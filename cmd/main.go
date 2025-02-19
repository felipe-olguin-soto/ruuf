package main

import (
	"fmt"
	"log"
	"ruuf/infrastructure/http"
)

func main() {
	fmt.Println("Starting application...")
	if err := http.StartServer(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

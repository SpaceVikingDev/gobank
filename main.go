package main

import (
	"fmt"
	"log"
)

func main() {
	configPath := "config.yml"
	config, err := readConfig(configPath)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Printf("Database host: %s\n", config.Database.Host)

	store, err := NewPostgresStore(config)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", store)

	// server := NewAPIServer(":3000", store)
	// server.Run()
}

package main

import (
	// golang package
	"log"

	// internal package
	"github.com/chaidiryahya/integration-test/app"
)

func main() {

	log.Println("[Integration Test] Starting")

	err := app.IntegrationTest()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("[Integration Test] End")
}

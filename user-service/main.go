package main

import (
	"log"
	"time"
	"user-service/cli"
)

func main() {
	// Ensure time zone can be loaded
	if _, err := time.LoadLocation("Africa/Tunis"); err != nil {
		log.Fatal(err)
	}

	// Start the CLI
	if err := cli.Execute(); err != nil {
		log.Fatal(err)
	}
}

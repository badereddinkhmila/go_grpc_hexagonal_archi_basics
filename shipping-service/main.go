package main

import (
	"log"
	"shipping-service/cli"
	"time"
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

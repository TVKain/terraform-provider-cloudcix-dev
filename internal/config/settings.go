package config

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Settings holds all configuration values for the CloudCIX SDK
type Settings struct {
	// Base API URL - service endpoints will be constructed from this
	CLOUDCIX_API_URL string
	// Legacy: kept for backward compatibility
	CLOUDCIX_API_V2_URL   string
	CLOUDCIX_API_VERSION  string
	CLOUDCIX_API_USERNAME string
	CLOUDCIX_API_PASSWORD string
	CLOUDCIX_API_KEY      string
	// Default region ID for creating resources
	CLOUDCIX_REGION_ID int
}

// DefaultSettings returns the default configuration
func DefaultSettings() *Settings {
	return &Settings{
		CLOUDCIX_API_URL:     "https://api.cloudcix.com/", // Base API URL
		CLOUDCIX_API_VERSION: "5.0",
	}
}

// LoadSettings loads configuration from environment variables and optionally from a file
func LoadSettings(settingsFile ...string) (*Settings, error) {
	settings := DefaultSettings()

	// Load from environment variables first
	if val := os.Getenv("CLOUDCIX_API_URL"); val != "" {
		settings.CLOUDCIX_API_URL = val
	}
	if val := os.Getenv("CLOUDCIX_API_V2_URL"); val != "" {
		settings.CLOUDCIX_API_V2_URL = val
	}
	if val := os.Getenv("CLOUDCIX_API_VERSION"); val != "" {
		settings.CLOUDCIX_API_VERSION = val
	}
	if val := os.Getenv("CLOUDCIX_API_USERNAME"); val != "" {
		settings.CLOUDCIX_API_USERNAME = val
	}
	if val := os.Getenv("CLOUDCIX_API_PASSWORD"); val != "" {
		settings.CLOUDCIX_API_PASSWORD = val
	}
	if val := os.Getenv("CLOUDCIX_API_KEY"); val != "" {
		settings.CLOUDCIX_API_KEY = val
	}
	if val := os.Getenv("CLOUDCIX_REGION_ID"); val != "" {
		if regionID, err := strconv.Atoi(val); err == nil {
			settings.CLOUDCIX_REGION_ID = regionID
		}
	}

	// If a settings file is provided, load from it (overrides env vars)
	if len(settingsFile) > 0 && settingsFile[0] != "" {
		err := settings.loadFromFile(settingsFile[0])
		if err != nil {
			return nil, fmt.Errorf("failed to load settings file: %w", err)
		}
	}

	return settings, nil
}

// loadFromFile loads settings from a file in KEY=VALUE format
func (s *Settings) loadFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Parse KEY=VALUE
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.Trim(strings.TrimSpace(parts[1]), `"`)

		switch key {
		case "CLOUDCIX_API_URL":
			s.CLOUDCIX_API_URL = value
		case "CLOUDCIX_API_V2_URL":
			s.CLOUDCIX_API_V2_URL = value
		case "CLOUDCIX_API_VERSION":
			s.CLOUDCIX_API_VERSION = value
		case "CLOUDCIX_API_USERNAME":
			s.CLOUDCIX_API_USERNAME = value
		case "CLOUDCIX_API_PASSWORD":
			s.CLOUDCIX_API_PASSWORD = value
		case "CLOUDCIX_API_KEY":
			s.CLOUDCIX_API_KEY = value
		case "CLOUDCIX_REGION_ID":
			if regionID, err := strconv.Atoi(value); err == nil {
				s.CLOUDCIX_REGION_ID = regionID
			}
		}
	}

	return scanner.Err()
}

// Validate checks if all required settings are present
func (s *Settings) Validate() error {
	if s.CLOUDCIX_API_USERNAME == "" {
		return fmt.Errorf("CLOUDCIX_API_USERNAME is required")
	}
	if s.CLOUDCIX_API_PASSWORD == "" {
		return fmt.Errorf("CLOUDCIX_API_PASSWORD is required")
	}
	if s.CLOUDCIX_API_KEY == "" {
		return fmt.Errorf("CLOUDCIX_API_KEY is required")
	}
	if s.CLOUDCIX_REGION_ID == 0 {
		return fmt.Errorf("CLOUDCIX_REGION_ID is required")
	}
	return nil
}

// MembershipURL returns the membership API URL
func (s *Settings) MembershipURL() string {
	// Use V2 URL if set (backward compatibility), otherwise construct from base
	if s.CLOUDCIX_API_V2_URL != "" {
		return s.CLOUDCIX_API_V2_URL
	}

	// For CloudCIX, we need to use the subdomain approach since api.cloudcix.com doesn't exist
	// Convert base URL to membership subdomain
	baseURL := strings.TrimSuffix(s.CLOUDCIX_API_URL, "/")
	if strings.Contains(baseURL, "api.cloudcix.com") {
		return "https://membership.api.cloudcix.com/"
	}
	// Fallback to appending /membership/ for other base URLs
	return baseURL + "/membership/"
}

// ComputeURL returns the compute API URL
func (s *Settings) ComputeURL() string {
	// For CloudCIX, we need to use the subdomain approach
	baseURL := strings.TrimSuffix(s.CLOUDCIX_API_URL, "/")
	if strings.Contains(baseURL, "api.cloudcix.com") {
		return "https://compute.api.cloudcix.com/"
	}
	// Fallback to appending /compute/ for other base URLs
	return baseURL + "/compute/"
}

// GetRegionID returns the configured region ID
func (s *Settings) GetRegionID() int {
	return s.CLOUDCIX_REGION_ID
}

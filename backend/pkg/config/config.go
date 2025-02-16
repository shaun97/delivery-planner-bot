package config

import (
    "fmt"
    "os"
    "strconv"

    "github.com/joho/godotenv"
)

type Config struct {
    Server struct {
        Port int
    }
    GoogleMaps struct {
        APIKey string
    }
}

func Load() (*Config, error) {
    // Load .env file if it exists
    if err := godotenv.Load(); err != nil {
        // Only log a warning as .env file is optional (might use environment variables directly)
        fmt.Printf("Warning: .env file not found: %v\n", err)
    }

    cfg := &Config{}

    // Server config
    port := os.Getenv("PORT")
    if port == "" {
        cfg.Server.Port = 8080 // Default port
    } else {
        p, err := strconv.Atoi(port)
        if err != nil {
            return nil, fmt.Errorf("invalid PORT value: %w", err)
        }
        cfg.Server.Port = p
    }

    // Google Maps config
    apiKey := os.Getenv("GOOGLE_MAPS_API_KEY")
    if apiKey == "" {
        return nil, fmt.Errorf("GOOGLE_MAPS_API_KEY environment variable is required")
    }
    cfg.GoogleMaps.APIKey = apiKey

    return cfg, nil
}

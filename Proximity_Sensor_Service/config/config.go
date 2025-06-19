package config

import (
    "log"
    "os"
    "strconv"
)

type Config struct {
    FirebaseCredPath string
    PollIntervalSec  int
}

func LoadConfig() *Config {
    firebasePath := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
    if firebasePath == "" {
        log.Fatal("GOOGLE_APPLICATION_CREDENTIALS not set in .env")
    }

    intervalStr := os.Getenv("POLL_INTERVAL_SEC")
    if intervalStr == "" {
        intervalStr = "5" // Default: 5 seconds
    }

    interval, err := strconv.Atoi(intervalStr)
    if err != nil || interval <= 0 {
        log.Fatalf("Invalid POLL_INTERVAL_SEC: %s", intervalStr)
    }

    return &Config{
        FirebaseCredPath: firebasePath,
        PollIntervalSec:  interval,
    }
}

package main

import (
    "log"
    "time"

    "github.com/yourusername/proximity-sensor-service/config"
    "github.com/yourusername/proximity-sensor-service/firebase"
    "github.com/yourusername/proximity-sensor-service/sensor"
    "github.com/joho/godotenv"
)

func main() {
    _ = godotenv.Load()

    cfg := config.LoadConfig()

    err := firebase.InitFirebase()
    if err != nil {
        log.Fatalf("Firebase init failed: %v", err)
    }

    s, err := sensor.InitSensor()
    if err != nil {
        log.Fatalf("Sensor init failed: %v", err)
    }

    for {
        distance, err := s.ReadDistance()
        if err != nil {
            log.Printf("Sensor read error: %v", err)
        } else {
            log.Printf("Distance: %.2f cm", distance)
            firebase.UploadDistance(distance)
        }

        time.Sleep(time.Duration(cfg.PollIntervalSec) * time.Second)
    }
}

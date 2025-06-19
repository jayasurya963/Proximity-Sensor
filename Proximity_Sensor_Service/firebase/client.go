package firebase

import (
    "context"
    "log"
    "os"

    firebase "firebase.google.com/go/v4"
    "firebase.google.com/go/v4/db"
    "google.golang.org/api/option"
)

var client *db.Client

func InitFirebase() error {
    opt := option.WithCredentialsFile(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))

    app, err := firebase.NewApp(context.Background(), nil, opt)
    if err != nil {
        return err
    }

    c, err := app.Database(context.Background())
    if err != nil {
        return err
    }

    client = c
    return nil
}

func UploadDistance(distance float64) {
    ref := client.NewRef("sensors/proximity")
    err := ref.Set(context.Background(), map[string]interface{}{
        "distance_cm": distance,
        "timestamp":   db.ServerValueTimestamp,
    })
    if err != nil {
        log.Printf("error uploading distance: %v", err)
    }
}


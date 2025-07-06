package Common

import (
    "context"
    "firebase.google.com/go/v4"
    "firebase.google.com/go/v4/auth"
    "google.golang.org/api/option"
    "log"
)

var FirebaseAuthClient *auth.Client

func InitFirebase() {
    app, err := firebase.NewApp(context.Background(), nil, option.WithCredentialsFile("serviceAccountKey.json"))
    if err != nil {
        log.Fatalf("Error initializing Firebase app: %v", err)
    }

    FirebaseAuthClient, err = app.Auth(context.Background())
    if err != nil {
        log.Fatalf("Error initializing Firebase Auth: %v", err)
    }
}

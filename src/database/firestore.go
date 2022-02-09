package models

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"

	"github.com/Real-Dev-Squad/gopher-cloud-service/src/config"
)

func createFirebaseApp(ctx context.Context) *firebase.App {
	// Use a service account
	sa := option.WithCredentialsJSON([]byte(config.Global.FirestoreCred))
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatalln(err)
	}
	return app
}

func CreateFireStoreClient(ctx context.Context) *firestore.Client {
	app := createFirebaseApp(ctx)
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	return client
}

package models

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/Real-Dev-Squad/website-learn-backend/src/config"
	"google.golang.org/api/option"
)

func createFirebaseApp(ctx context.Context) *firebase.App {
	sa := option.WithCredentialsJSON([]byte(config.Global.FirestoreCred))
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		log.Fatal(err)
	}
	return app
}

func createFirestoreClient(ctx context.Context) *firestore.Client {
	app := createFirebaseApp(ctx)
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func GetCollection(name string) *firestore.CollectionRef {
	client := createFirestoreClient(context.Background())
	model := client.Collection(name)
	return model
}

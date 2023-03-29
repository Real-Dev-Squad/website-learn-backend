package database

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"

	"github.com/Real-Dev-Squad/gopher-cloud-service/src/config"
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

func GetCollection(collectionName string) *firestore.CollectionRef {
	client := createFirestoreClient(context.Background())
	model := client.Collection(collectionName)
	return model
}

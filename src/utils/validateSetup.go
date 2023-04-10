package utils

import (
	"context"
	"log"

	"github.com/Real-Dev-Squad/gopher-cloud-service/src/models"
	"google.golang.org/api/iterator"
)

func ValidateSetup() {

	testModel := models.GetCollection("test")

	data := map[string]interface{}{"Name": "test"}

	_, _, err := testModel.Add(context.Background(), data)

	query := testModel.Where("Name", "==", "test")

	iter := query.Documents(context.Background())
	defer iter.Stop()

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		doc.Ref.Delete(context.Background())
	}

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("✅ Code has the permission to write to the database")
		log.Println("✅ Code has the permission to read from the database")
		log.Println("🎉 Congratulations the firestore setup is complete")
	}

}

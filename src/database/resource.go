package models

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	firestoreConstant "github.com/Real-Dev-Squad/gopher-cloud-service/src/constants/firestore"
	resourceSchema "github.com/Real-Dev-Squad/gopher-cloud-service/src/schema/resource"
	"google.golang.org/api/iterator"
)

const RESOURCE_COLLECTION = firestoreConstant.COLLECTION_RESOURCE

func GetResources(limit int, startAfter string) (*[]resourceSchema.Resources, error) {
	ctx := context.Background()
	var firestoreClient = CreateFireStoreClient(ctx)
	defer firestoreClient.Close()

	query := firestoreClient.Collection(RESOURCE_COLLECTION).OrderBy("publishedDate", firestore.Desc).Limit(limit)

	if startAfter != "" {
		startAfterDoc, _ := firestoreClient.Collection(RESOURCE_COLLECTION).Doc(startAfter).Get(ctx)
		query = query.StartAfter(startAfterDoc)
	}

	iter := query.Documents(ctx)

	resources := []resourceSchema.Resources{}
	var resource resourceSchema.Resources

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		doc.DataTo(&resource)
		resources = append(resources, resource)
		fmt.Print("1")
		if err != nil {
			log.Printf("Failed getting resources: %v", err)
			return nil, err
		}
	}

	return &resources, nil
}

func GetResource(id string) (*resourceSchema.Resources, error) {
	ctx := context.Background()
	var firestoreClient = CreateFireStoreClient(ctx)
	defer firestoreClient.Close()

	resourceDocument, err := firestoreClient.Collection(RESOURCE_COLLECTION).Doc(id).Get(ctx)
	if err != nil {
		log.Printf("Failed getting resource: %s \n %v", id, err)
		return nil, err
	}

	var resource resourceSchema.Resources
	err = resourceDocument.DataTo(&resource)
	if err != nil {
		log.Printf("Failed parsing resource to resource schema: %s \n %v", id, err)
		return nil, err
	}

	return &resource, nil
}

func PostResource(resource resourceSchema.Resources) (*resourceSchema.Resources, error) {
	ctx := context.Background()
	var firestoreClient = CreateFireStoreClient(ctx)
	defer firestoreClient.Close()

	ref := firestoreClient.Collection(RESOURCE_COLLECTION).NewDoc()

	resource.Id = ref.ID

	_, err := ref.Set(ctx, resource)

	if err != nil {
		log.Printf("Failed adding resource: %v \n %v", resource, err)
		return nil, err
	}
	return &resource, nil
}

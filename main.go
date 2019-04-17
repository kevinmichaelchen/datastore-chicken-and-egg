package main

import (
	"cloud.google.com/go/datastore"
	"context"
	"log"
)

const (
	PermissionsEntity = "Permissions"
	FolderEntity      = "Folder"
)

type Folder struct {
	ParentID string
}

func createFolder(ctx context.Context, client *datastore.Client, key *datastore.Key) {
	log.Println("Creating key:", key.String())
	f := &Folder{}
	parentKey := key.Parent
	if parentKey != nil {
		f.ParentID = parentKey.Name
	}
	_, err := client.Put(ctx, key, f)
	if err != nil {
		log.Fatalf("Error creating folder: %s", err.Error())
	}
}

func main() {
	ctx := context.TODO()

	// Read Datastore Project ID
	log.Println("Reading configs")
	config := LoadConfig()

	// Create client
	log.Println("Creating client")
	client, err := datastore.NewClient(ctx, config.DatastoreProjectID)
	if err != nil {
		log.Fatalf("Error creating client. %s", err.Error())
	}

	// Create a root folder
	sportsKey := datastore.NameKey(FolderEntity, "sports", nil)
	createFolder(ctx, client, sportsKey)

	contactSportsKey := datastore.NameKey(FolderEntity, "contact-sports", sportsKey)
	createFolder(ctx, client, contactSportsKey)

	winterSportsKey := datastore.NameKey(FolderEntity, "winter-sports", sportsKey)
	createFolder(ctx, client, winterSportsKey)

	iceHockeyKey := datastore.NameKey(FolderEntity, "ice-hockey", winterSportsKey)
	createFolder(ctx, client, iceHockeyKey)

	snowboardingKey := datastore.NameKey(FolderEntity, "snowboarding", winterSportsKey)
	createFolder(ctx, client, snowboardingKey)

	trickKey := datastore.NameKey(FolderEntity, "trick-snowboarding", snowboardingKey)
	createFolder(ctx, client, trickKey)

	// Let's do a global query for all entities w/ NodeID = 1
	q := datastore.NewQuery(FolderEntity).
		KeysOnly().
		Ancestor(winterSportsKey).
		Filter("ParentID =", winterSportsKey.Name)

	log.Println("Doing query")
	keys, err := client.GetAll(ctx, q, nil)
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

	for i, key := range keys {
		log.Printf("key #%d = %s\n", i, key.String())
	}
}

package main

import (
	"cloud.google.com/go/datastore"
	"context"
	"fmt"
	"log"
)

const FolderEntity = "Folder"

type Folder struct {
	ID       string
	ParentID string
}

func (f Folder) String() string {
	return fmt.Sprintf("Folder[ID=%s, ParentID=%s]", f.ID, f.ParentID)
}

func main() {
	ctx := context.TODO()

	// Read Datastore Project ID
	config := LoadConfig()

	// Create client
	client, err := datastore.NewClient(ctx, config.DatastoreProjectID)
	if err != nil {
		log.Fatalf("Error creating client. %s", err.Error())
	}

	// Let's make a root folder!!!
	rootFolderID := "1"
	rootFolderKey := datastore.NameKey(FolderEntity, rootFolderID, nil)
	rootFolder := &Folder{
		ID:       rootFolderID,
		ParentID: "",
	}
	_, err = client.Put(ctx, rootFolderKey, rootFolder)
	if err != nil {
		log.Fatalf("Error putting root folder. %s", err.Error())
	}

	// Let's make a subfolder under
	// the root folder we just created!!!
	subFolderID := "2"
	subFolderKey := datastore.NameKey(FolderEntity, subFolderID, rootFolderKey)
	subFolder := &Folder{
		ID:       subFolderID,
		ParentID: rootFolderID,
	}
	_, err = client.Put(ctx, subFolderKey, subFolder)
	if err != nil {
		log.Fatalf("Error putting subfolder. %s", err.Error())
	}

	// OKAY!
	// Let's do a DB lookup
	//badKey := datastore.NameKey(FolderEntity, subFolderID, nil)
	goodKey := datastore.NameKey(FolderEntity, subFolderID, rootFolderKey)
	var f Folder
	err = client.Get(ctx, goodKey, &f)
	if err != nil {
		log.Fatalf("Error reading: %s", err.Error())
	}

	log.Println("GET RESULT =", f.String())
}

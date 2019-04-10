package main

import (
	"cloud.google.com/go/datastore"
	"context"
	"fmt"
	"log"
)

const (
	PermissionsEntity = "Permissions"
)

type Permissions struct {
	NodeID string
}

func (p Permissions) String() string {
	return fmt.Sprintf("Permissions[NodeID=%s]", p.NodeID)
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


	// Grant User 5 permissions on Node 1
	n1 := "1"
	k1 := datastore.NameKey(PermissionsEntity, n1, datastore.NameKey("User", "5", nil))
	_, err = client.Put(
		ctx,
		k1,
		&Permissions{NodeID: n1})
	if err != nil {
		log.Fatalf("Error creating permissions: %s", err.Error())
	}

	// Grant Team 6 permissions on Node 1
	n2 := "45"
	k2 := datastore.NameKey(PermissionsEntity, n2, datastore.NameKey("Team", "6", nil))
	_, err = client.Put(
		ctx,
		k2,
		&Permissions{NodeID: n2})
	if err != nil {
		log.Fatalf("Error creating permissions: %s", err.Error())
	}

	// OKAY!
	// Let's do a DB lookup
	//badKey := datastore.NameKey(FolderEntity, subFolderID, nil)
	var p1, p2 Permissions
	if err := client.Get(ctx, k1, &p1); err != nil {
		log.Fatalf("Error reading: %s", err.Error())
	}
	if err := client.Get(ctx, k2, &p2); err != nil {
		log.Fatalf("Error reading: %s", err.Error())
	}

	log.Println("p1 =", p1.String())
	log.Println("p2 =", p2.String())
}

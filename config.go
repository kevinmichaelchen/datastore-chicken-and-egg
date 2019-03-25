package main

import (
	"encoding/json"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
)

const (
	flagForGoogleApplicationCredentials = "google_application_credentials"
	flagForDatastoreProjectID           = "datastore_project_id"
)

type Config struct {
	GoogleApplicationCredentials string
	DatastoreProjectID           string
}

func (c Config) String() string {
	b, err := json.Marshal(c)
	if err != nil {
		log.Fatalf("Could not marshal config to string: %v", err)
	}
	return string(b)
}

func LoadConfig() Config {
	c := Config{}

	flag.String(flagForGoogleApplicationCredentials, c.GoogleApplicationCredentials, "Absolute path to Google Application Credentials")
	flag.String(flagForDatastoreProjectID, c.DatastoreProjectID, "Google Cloud Datastore Project ID")

	flag.Parse()

	viper.BindPFlag(flagForGoogleApplicationCredentials, flag.Lookup(flagForGoogleApplicationCredentials))
	viper.BindPFlag(flagForDatastoreProjectID, flag.Lookup(flagForDatastoreProjectID))

	viper.AutomaticEnv()

	c.GoogleApplicationCredentials = viper.GetString(flagForGoogleApplicationCredentials)
	c.DatastoreProjectID = viper.GetString(flagForDatastoreProjectID)

	return c
}
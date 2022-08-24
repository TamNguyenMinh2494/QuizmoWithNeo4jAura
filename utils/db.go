package utils

import (
	"log"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type Database struct {
	Client neo4j.Driver
}

func NewDatabase(config *Config) *Database {
	driver, err := neo4j.NewDriver(config.Uri, neo4j.BasicAuth(config.Username, config.Password, ""))
	if err != nil {
		log.Fatalf("Failed to create driver: %s", err)
	}
	// defer driver.Close()
	return &Database{
		Client: driver,
	}
}

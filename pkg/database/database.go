package database

import (
	"os"

	log "github.com/sirupsen/logrus"
)

// Initialize the database
func (db *Database) Initialize() {
	parkFilename := "./data/parks.csv"
	speciesFilename := "./data/species.csv"

	parkFile, err := os.Open(parkFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer parkFile.Close()

	speciesFile, err := os.Open(speciesFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer speciesFile.Close()

	db.parseParkFile(parkFile)
	db.parseSpeciesFile(speciesFile)
}

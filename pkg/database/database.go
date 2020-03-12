package database

import (
	"errors"
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

// GetParkByID - Fetch park struct by id
func (db *Database) GetParkByID(id string) (park *Park, err error) {
	park, ok := (*db.Parks)[id]
	if ok {
		return park, nil
	}
	return nil, errors.New("park not found")
}

// GetParks - Fetch all parks
func (db *Database) GetParks() *[]*Park {
	parks := make([]*Park, 0, len(*db.Parks))

	for _, value := range *db.Parks {
		parks = append(parks, value)
	}
	return &parks
}

// GetSpeciesByID - Fetch species struct by id
func (db *Database) GetSpeciesByID(id string) (species *Species, err error) {
	species, ok := (*db.Species)[id]
	if ok {
		return species, nil
	}
	return nil, errors.New("species not found")
}

// GetSpeciesByPark - Get all the species in a park
func (db *Database) GetSpeciesByPark(parkID string) *[]*Species {
	ret := make([]*Species, 0, 0)
	for _, item := range *db.ParkRecords {
		if item.Park.ID == parkID {
			ret = append(ret, item.Species)
		}
	}

	return &ret
}

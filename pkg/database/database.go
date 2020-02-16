package database

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

// Database is the entire data object, parsed from csv files
type Database struct {
	Parks       *map[string]*Park
	Species     *map[string]*Species
	ParkRecords *[]*ParkRecord
}

// Park is a national park
type Park struct {
	ID        string   `json:"id"`
	Name      string   `json:"park_name"`
	States    []string `json:"states"`
	Acres     int      `json:"acres"`
	Latitude  float64  `json:"latitude"`
	Longitude float64  `json:"longitude"`
}

// Species is an individual species
type Species struct {
	ID                 string   `json:"id"`
	Category           string   `json:"category"`
	Order              string   `json:"order"`
	Family             string   `json:"family"`
	ScientificName     string   `json:"scientific_name"`
	CommonNames        []string `json:"common_names"`
	ConservationStatus string   `json:"conservation_status,omitempty"`
}

// ParkRecord is a species that has been recorded in a national park
type ParkRecord struct {
	ID           string   `json:"id"`
	Park         *Park    `json:"park_id"`
	Species      *Species `json:"species_id"`
	RecordStatus string   `json:"record_status"`
	Occurrence   string   `json:"occurrence"`
	Nativeness   string   `json:"nativeness,omitempty"`
	Abundance    string   `json:"abundance,omitempty"`
	Seasonality  string   `json:"seasonality,omitempty"`
}

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
}

func (db *Database) parseParkFile(r io.Reader) {
	reader := csv.NewReader(r)

	ret := make(map[string]*Park)

	for {
		row, err := reader.Read()
		if err == io.EOF {
			log.Debug("end of file")
			break
		} else if err != nil {
			log.Error(err)
			break
		}

		if row[0] == "Park Code" {
			continue
		}

		states := strings.Split(row[2], ",")
		acres, _ := strconv.Atoi(row[3])
		latitude, _ := strconv.ParseFloat(row[4], 64)
		longitude, _ := strconv.ParseFloat(row[5], 64)

		park := &Park{
			ID:        row[0],
			Name:      row[1],
			States:    states,
			Acres:     acres,
			Latitude:  latitude,
			Longitude: longitude,
		}

		ret[row[0]] = park
	}
	db.Parks = &ret
}

func parseSpeciesFile(r io.Reader) *[]*Species {
	reader := csv.NewReader(r)

	ret := make([]*Species, 0, 0)

	for {
		row, err := reader.Read()
		if err == io.EOF {
			log.Debug("end of file")
			break
		} else if err != nil {
			log.Error(err)
			break
		}

		if row[0] == "Species ID" {
			continue
		}

		commonNames := strings.Split(row[6], ",")

		species := &Species{
			ID:                 row[0],
			Category:           row[2],
			Order:              row[3],
			Family:             row[4],
			ScientificName:     row[5],
			CommonNames:        commonNames,
			ConservationStatus: row[12],
		}

		ret = append(ret, species)
	}
	return &ret
}

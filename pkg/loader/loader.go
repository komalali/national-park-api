package loader

import (
	"encoding/csv"
	"io"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

type Database struct {
	Parks   *[]*ParkData
	Species *[]*SpeciesData
}

// ParkData is data about a national park
type ParkData struct {
	ID        string   `json:"park_id"`
	Name      string   `json:"park_name"`
	States    []string `json:"states"`
	Acres     int      `json:"acres"`
	Latitude  float64  `json:"latitude"`
	Longitude float64  `json:"longitude"`
}

// SpeciesData is data about species found in national parks
type SpeciesData struct {
	ID                 string   `json:"species_id"`
	ParkName           string   `json:"park_name"`
	Category           string   `json:"category"`
	Order              string   `json:"order"`
	Family             string   `json:"family"`
	ScientificName     string   `json:"scientific_name"`
	CommonNames        []string `json:"common_names"`
	RecordStatus       string   `json:"record_status"`
	Occurrence         string   `json:"occurrence"`
	Nativeness         string   `json:"nativeness,omitempty"`
	Abundance          string   `json:"abundance,omitempty"`
	Seasonality        string   `json:"seasonality,omitempty"`
	ConservationStatus string   `json:"conservation_status,omitempty"`
}

// LoadData loads the data from files into memory
func LoadData(parkReader io.Reader, speciesReader io.Reader) *Database {
	parks := loadParkData(parkReader)
	species := loadSpeciesData(speciesReader)

	return &Database{
		Parks:   parks,
		Species: species,
	}
}

// loadParkData loads the park csv data into memory
func loadParkData(r io.Reader) *[]*ParkData {
	reader := csv.NewReader(r)

	ret := make([]*ParkData, 0, 0)

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

		park := &ParkData{
			ID:        row[0],
			Name:      row[1],
			States:    states,
			Acres:     acres,
			Latitude:  latitude,
			Longitude: longitude,
		}

		ret = append(ret, park)
	}
	return &ret
}

// loadSpeciesData loads the species csv data into memory
func loadSpeciesData(r io.Reader) *[]*SpeciesData {
	reader := csv.NewReader(r)

	ret := make([]*SpeciesData, 0, 0)

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

		species := &SpeciesData{
			ID:                 row[0],
			ParkName:           row[1],
			Category:           row[2],
			Order:              row[3],
			Family:             row[4],
			ScientificName:     row[5],
			CommonNames:        commonNames,
			RecordStatus:       row[7],
			Occurrence:         row[8],
			Nativeness:         row[9],
			Abundance:          row[10],
			Seasonality:        row[11],
			ConservationStatus: row[12],
		}

		ret = append(ret, species)
	}
	return &ret
}

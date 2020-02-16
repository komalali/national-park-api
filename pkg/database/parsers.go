package database

import (
	"encoding/csv"
	"io"
	"strconv"
	"strings"

	"github.com/gosimple/slug"
	log "github.com/sirupsen/logrus"
)

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
		for i := range states {
			states[i] = strings.TrimSpace(states[i])
		}
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

func (db *Database) parseSpeciesFile(r io.Reader) {
	reader := csv.NewReader(r)

	speciesMap := make(map[string]*Species)
	parkRecords := make([]*ParkRecord, 0, 0)

	for {
		var species *Species

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

		parkID := strings.Split(row[0], "-")[0]
		park := (*db.Parks)[parkID]
		scientificNameSlug := slug.Make(row[5])

		foundSpecies, ok := speciesMap[scientificNameSlug]
		if !ok {
			commonNames := strings.Split(row[6], ",")
			for i := range commonNames {
				commonNames[i] = strings.TrimSpace(commonNames[i])
			}
			species = &Species{
				ID:                 scientificNameSlug,
				Category:           row[2],
				Order:              row[3],
				Family:             row[4],
				ScientificName:     row[5],
				CommonNames:        commonNames,
				ConservationStatus: row[12],
				Parks:              []string{},
			}
			species.addPark(park)
			speciesMap[scientificNameSlug] = species
		} else {
			foundSpecies.addPark(park)
		}

		record := &ParkRecord{
			ID:           row[0],
			Park:         park,
			Species:      species,
			RecordStatus: row[7],
			Occurrence:   row[8],
			Nativeness:   row[9],
			Abundance:    row[10],
			Seasonality:  row[11],
		}
		parkRecords = append(parkRecords, record)
	}

	db.ParkRecords = &parkRecords
	db.Species = &speciesMap
}

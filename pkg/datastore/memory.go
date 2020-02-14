package datastore

import (
	"os"
	"strings"

	"github.com/komalali/national-parks/api/pkg/loader"
	log "github.com/sirupsen/logrus"
)

// Parks - collection of parks
type Data struct {
	Store *loader.Database `json:"store"`
}

// Initialize the Park store
func (d *Data) Initialize() {
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

	d.Store = loader.LoadData(parkFile, speciesFile)
}

// GetParkById - Find park information by code
func (d *Data) GetParkByID(id string) *loader.ParkData {
	parks := d.Store.Parks
	ret := Filter(parks, func(v *loader.ParkData) bool {
		return strings.ToLower(v.ID) == strings.ToLower(id)
	})
	if len(*ret) > 0 {
		return (*ret)[0]
	}
	return nil
}

// GetAllParks - Return all the parks
func (d *Data) GetAllParks() *[]*loader.ParkData {
	return d.Store.Parks
}

// Filter - returns a slice of parks
func Filter(vs *[]*loader.ParkData, f func(*loader.ParkData) bool) *[]*loader.ParkData {
	vsf := make([]*loader.ParkData, 0)
	for _, v := range *vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return &vsf
}

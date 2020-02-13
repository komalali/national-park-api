package datastore

import (
	"log"
	"os"
	"strings"

	"github.com/komalali/national-parks/api/pkg/loader"
)

// Parks - collection of parks
type Parks struct {
	Store *[]*loader.ParkData `json:"store"`
}

// Initialize the Park store
func (p *Parks) Initialize() {
	filename := "./data/parks.csv"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	p.Store = loader.LoadParkData(file)
}

// SearchByID - Find park information by code
func (p *Parks) SearchByID(id string) *loader.ParkData {
	ret := Filter(p.Store, func(v *loader.ParkData) bool {
		return strings.ToLower(v.ID) == strings.ToLower(id)
	})
	if len(*ret) > 0 {
		return (*ret)[0]
	}
	return nil
}

func Filter(vs *[]*loader.ParkData, f func(*loader.ParkData) bool) *[]*loader.ParkData {
	vsf := make([]*loader.ParkData, 0)
	for _, v := range *vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return &vsf
}

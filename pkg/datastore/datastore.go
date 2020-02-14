package datastore

import "github.com/komalali/national-parks/api/pkg/loader"

// ParkStore - Data store for parks
type DataStore interface {
	Initialize()
	GetParkByID(id string) *loader.ParkData
	GetAllParks() *[]*loader.ParkData
}

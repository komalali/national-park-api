package datastore

import "github.com/komalali/national-parks/api/pkg/loader"

// ParkStore - Data store for parks
type ParkStore interface {
	Initialize()
	SearchByID(code string) *loader.ParkData
	GetAllParks() *[]*loader.ParkData
}

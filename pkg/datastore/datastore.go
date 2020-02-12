package datastore

import "github.com/komalali/national-parks/api/pkg/loader"

// ParkStore - Data store for parks
type ParkStore interface {
	Initialize()
	SearchByCode(code string) *loader.ParkData
}

package datastore

import "github.com/komalali/national-parks/api/pkg/loader"

// Store - Data store for parks
type Store interface {
	Initialize()
	SearchByID(code string) *loader.ParkData
}

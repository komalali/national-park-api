package datastore

import "github.com/komalali/national-parks/api/pkg/loader"

// Store - Data store for parks
type Store interface {
	Initialize()
	SearchByCode(code string) *loader.ParkData
}

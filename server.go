package main

import (
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/komalali/national-parks/api/pkg/database"
	log "github.com/sirupsen/logrus"
)

var db database.Database

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)

	log.Infof("%s took %s", name, elapsed)
}

func init() {
	defer timeTrack(time.Now(), "file load")
	db.Initialize()
}

func main() {
	species, err := db.GetSpeciesByID("colpophyllia-natans")
	if err != nil {
		log.Error(err)
	} else {
		spew.Dump(species)
	}

	park, err := db.GetParkByID("ARCH")
	if err != nil {
		log.Error(err)
	} else {
		spew.Dump(park)
	}

	// parks := db.GetParks()
	// spew.Dump(len(*parks))

	speciesAtArches := db.GetSpeciesByPark("ARCH")
	spew.Dump(speciesAtArches)

	// router := mux.NewRouter()

	// api := router.PathPrefix("/api/v1").Subrouter()

	// api.HandleFunc("", home).Methods(http.MethodGet)
	// api.HandleFunc("/parks/id/{id}", getParkByID).Methods(http.MethodGet)
	// api.HandleFunc("/parks", getAllParks).Methods(http.MethodGet)

	// log.Fatal(http.ListenAndServe(":8080", router))
}

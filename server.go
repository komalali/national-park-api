package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/komalali/national-parks/api/pkg/datastore"
	log "github.com/sirupsen/logrus"
)

var data datastore.DataStore

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)

	log.Infof("%s took %s", name, elapsed)
}

func init() {
	defer timeTrack(time.Now(), "file load")
	data = &datastore.Data{}
	data.Initialize()
}

func main() {
	router := mux.NewRouter()

	api := router.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("", home).Methods(http.MethodGet)
	api.HandleFunc("/parks/id/{id}", getParkByID).Methods(http.MethodGet)
	api.HandleFunc("/parks", getAllParks).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", router))
}

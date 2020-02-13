package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/komalali/national-parks/api/pkg/datastore"
	log "github.com/sirupsen/logrus"
)

var parks datastore.Store

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)

	log.Infof("%s took %s", name, elapsed)
}

func init() {
	defer timeTrack(time.Now(), "file load")
	parks = &datastore.Parks{}
	parks.Initialize()
}

func main() {
	router := mux.NewRouter()

	api := router.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "ok", "message": "api v1"}`))
	})
	api.HandleFunc("/parks/id/{id}", getParkByID).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", router))
}

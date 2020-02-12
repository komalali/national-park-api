package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/komalali/national-parks/api/pkg/datastore"
)

var parks datastore.Store

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func init() {
	defer timeTrack(time.Now(), "file load")
	parks = &datastore.Parks{}
	parks.Initialize()
}

func searchByID(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)

	w.Header().Set("Content-Type", "application/json")
	if val, ok := pathParams["id"]; ok {
		log.Printf("GET - searchByID called - id: %s", val)
		data := parks.SearchByCode(val)
		if data != nil {
			b, err := json.Marshal(data)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`{"error": "error marshalling data"}`))
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write(b)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "invalid park id"}`))
	}
}

func main() {
	router := mux.NewRouter()

	api := router.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "ok", "message": "api v1"}`))
	})
	api.HandleFunc("/parks/id/{id}", searchByID).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", router))
}

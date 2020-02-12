package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/komalali/national-parks/api/pkg/datastore"
)

var parks datastore.ParkStore

func searchByID(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)

	w.Header().Set("Content-Type", "application/json")
	if val, ok := pathParams["id"]; ok {
		log.Printf(val)
		data := parks.SearchByCode(val)
		log.Print(data)
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
	}
}

func main() {
	// router := mux.NewRouter()

	// api := router.PathPrefix("/api/v1").Subrouter()
	// api.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.WriteHeader(http.StatusOK)
	// 	w.Write([]byte(`{"status": "ok", "message": "api v1"}`))
	// })

	// api.HandleFunc("/parks/id/{id}", searchByID).Methods(http.MethodGet)

	// log.Fatal(http.ListenAndServe(":8080", router))
}

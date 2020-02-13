package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func getParkByID(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

	if val, ok := pathParams["id"]; ok {
		logger := log.WithFields(log.Fields{"method": "GET", "id": val})
		data := parks.SearchByID(val)
		if data != nil {
			b, err := json.Marshal(data)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`{"error": "error marshalling data"}`))
				logger.Errorf("error marshalling data")
				return
			}
			logger.Infof("searchByID called")
			w.WriteHeader(http.StatusOK)
			w.Write(b)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "invalid park id"}`))
		logger.Errorf("invalid park id")
	}
}

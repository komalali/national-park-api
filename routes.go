// package main

// import (
// 	"encoding/json"
// 	"net/http"

// 	"github.com/gorilla/mux"
// 	log "github.com/sirupsen/logrus"
// )

// func home(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte(`{"status": "ok", "message": "api v1"}`))
// 	log.WithFields(log.Fields{"method": "GET"}).Info("home called")
// }

// func getParkByID(w http.ResponseWriter, r *http.Request) {
// 	pathParams := mux.Vars(r)
// 	w.Header().Set("Content-Type", "application/json")

// 	if val, ok := pathParams["id"]; ok {
// 		logger := log.WithFields(log.Fields{"method": "GET", "id": val})
// 		result := data.GetParkByID(val)
// 		if result != nil {
// 			b, err := json.Marshal(result)
// 			if err != nil {
// 				w.WriteHeader(http.StatusInternalServerError)
// 				w.Write([]byte(`{"error": "error marshalling data"}`))
// 				logger.Errorf("error marshalling data")
// 				return
// 			}
// 			logger.Infof("searchByID called")
// 			w.WriteHeader(http.StatusOK)
// 			w.Write(b)
// 			return
// 		}
// 		w.WriteHeader(http.StatusBadRequest)
// 		w.Write([]byte(`{"error": "invalid park id"}`))
// 		logger.Errorf("invalid park id")
// 	}
// }

// func getAllParks(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	logger := log.WithFields(log.Fields{"method": "GET"})
// 	result := data.GetAllParks()

// 	b, err := json.Marshal(result)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte(`{"error": "error marshalling data"}`))
// 		logger.Errorf("error marshalling data")
// 		return
// 	}

// 	logger.Infof("getAllParks called")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(b)
// 	return
// }

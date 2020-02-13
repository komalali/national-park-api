package loader

import (
	"encoding/csv"
	"io"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

// ParkData is data about a national park
type ParkData struct {
	ID        string   `json:"park_id"`
	Name      string   `json:"park_name"`
	States    []string `json:"states"`
	Acres     int      `json:"acres"`
	Latitude  float64  `json:"latitude"`
	Longitude float64  `json:"longitude"`
}

// LoadParkData loads the csv data into memory
func LoadParkData(r io.Reader) *[]*ParkData {
	reader := csv.NewReader(r)

	ret := make([]*ParkData, 0, 0)

	for {
		row, err := reader.Read()
		if err == io.EOF {
			log.Debug("end of file")
			break
		} else if err != nil {
			log.Error(err)
			break
		}

		if row[0] == "Park Code" {
			continue
		}

		states := strings.Split(row[2], ",")
		acres, _ := strconv.Atoi(row[3])
		latitude, _ := strconv.ParseFloat(row[4], 64)
		longitude, _ := strconv.ParseFloat(row[5], 64)

		park := &ParkData{
			ID:        row[0],
			Name:      row[1],
			States:    states,
			Acres:     acres,
			Latitude:  latitude,
			Longitude: longitude,
		}

		ret = append(ret, park)
	}
	return &ret
}

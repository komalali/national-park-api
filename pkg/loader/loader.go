package loader

import (
	"encoding/csv"
	"io"
	"log"
	"strconv"
	"strings"
)

// ParkData is data about a national park
type ParkData struct {
	Code      string   `json:"Park Code"`
	Name      string   `json:"Park Name"`
	States    []string `json:"State"`
	Acres     int      `json:"Acres"`
	Latitude  float64  `json:"Latitude"`
	Longitude float64  `json:"Longitude"`
}

func LoadParkData(r io.Reader) *[]*ParkData {
	reader := csv.NewReader(r)

	ret := make([]*ParkData, 0, 0)

	for {
		row, err := reader.Read()
		if err == io.EOF {
			log.Println("End of File")
			break
		} else if err != nil {
			log.Println(err)
			break
		}

		states := strings.Split(row[2], ",")
		acres, _ := strconv.Atoi(row[3])
		latitude, _ := strconv.ParseFloat(row[4], 64)
		longitude, _ := strconv.ParseFloat(row[5], 64)

		park := &ParkData{
			Code:      row[0],
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

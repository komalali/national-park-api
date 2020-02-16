package database

// Database is the entire data object, parsed from csv files
type Database struct {
	Parks       *map[string]*Park
	Species     *map[string]*Species
	ParkRecords *[]*ParkRecord
}

// Park is a national park
type Park struct {
	ID        string   `json:"id"`
	Name      string   `json:"park_name"`
	States    []string `json:"states"`
	Acres     int      `json:"acres"`
	Latitude  float64  `json:"latitude"`
	Longitude float64  `json:"longitude"`
}

// Species is an individual species
type Species struct {
	ID                 string   `json:"id"`
	Category           string   `json:"category"`
	Order              string   `json:"order"`
	Family             string   `json:"family"`
	ScientificName     string   `json:"scientific_name"`
	CommonNames        []string `json:"common_names"`
	ConservationStatus string   `json:"conservation_status,omitempty"`
}

// ParkRecord is a species that has been recorded in a national park
type ParkRecord struct {
	ID           string   `json:"id"`
	Park         *Park    `json:"park_id"`
	Species      *Species `json:"species_id"`
	RecordStatus string   `json:"record_status"`
	Occurrence   string   `json:"occurrence"`
	Nativeness   string   `json:"nativeness,omitempty"`
	Abundance    string   `json:"abundance,omitempty"`
	Seasonality  string   `json:"seasonality,omitempty"`
}

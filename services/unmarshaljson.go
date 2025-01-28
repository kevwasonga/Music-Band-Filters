package services

import (
	"encoding/json"
	"fmt"
	"os"
)

type Band struct {
	Name             string              `json:"name"`
	FormationYear    int                 `json:"formationYear"`
	FirstAlbumYear   int                 `json:"firstAlbumYear"`
	Members          int                 `json:"members"`
	ConcertLocations map[string][]string `json:"concertLocations"` // Updated to match JSON structure
}

func LoadBands(filepath string) ([]Band, error) {
	const defaultFilePath = "data/bands.json"

	if filepath == "" {
		filepath = defaultFilePath
	}

	// Read the JSON file
	file, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("error reading JSON file: %w", err)
	}

	// Unmarshal the JSON data into a slice of Band structs
	var bands []Band
	if err := json.Unmarshal(file, &bands); err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	// Print the loaded bands for testing
	fmt.Println("Loaded bands:", bands)

	return bands, nil
}

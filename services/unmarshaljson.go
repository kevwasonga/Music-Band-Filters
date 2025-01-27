package services

import (
	"encoding/json"
	"fmt"
	"os"
)

type Band struct {
	Name             string   `json:"name"`             // Matches "name" in JSON
	FormationYear    int      `json:"formationYear"`    // Matches "formationYear" in JSON
	FirstAlbumYear   int      `json:"firstAlbumYear"`   // Matches "firstAlbumYear" in JSON
	Members          int      `json:"members"`          // Matches "members" in JSON
	ConcertLocations []string `json:"concertLocations"` // Matches "concertLocations" in JSON
}

func LoadBands(filepath string) ([]Band, error) {
	JSONFilePath := "data/bands.json" // Default file path for the JSON file

	if filepath == "" {
		filepath = JSONFilePath
	}

	// Read the JSON file
	file, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("error reading JSON file: %w", err)
	}

	// Unmarshal the JSON data into a slice of Band structs
	var bands []Band

	err = json.Unmarshal(file, &bands)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	// Print the loaded bands for testing
	fmt.Println("Loaded bands:", bands)

	return bands, nil
}

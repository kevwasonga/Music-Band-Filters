package models

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"fiter/services"
)

// FilterBandsHandler handles filtering of bands based on criteria
func FilterBandsHandler(w http.ResponseWriter, r *http.Request) {
	// Load bands from the service
	bands, err := services.LoadBands("")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Parse filter criteria from query parameters
	query := r.URL.Query()

	minFormationYear, _ := strconv.Atoi(query.Get("minFormationYear"))
	maxFormationYear, _ := strconv.Atoi(query.Get("maxFormationYear"))

	minFirstAlbumYear, _ := strconv.Atoi(query.Get("minFirstAlbumYear"))
	maxFirstAlbumYear, _ := strconv.Atoi(query.Get("maxFirstAlbumYear"))


	minMembers, _ := strconv.Atoi(query.Get("minMembers"))

	maxMembers, _ := strconv.Atoi(query.Get("maxMembers"))

	locations := query["locations"] // Get array of locations

	fmt.Println("minFrYear: ", minFirstAlbumYear,"\n maxFrYear: ", maxFirstAlbumYear,"\nminFoYear: ", minFormationYear,"\n maxFoYear: ", maxFormationYear,"\nminMembers: ", minMembers,"\nmaxMembers: ", maxMembers, "\nLocations: ", locations)

	// Apply filters to the bands
	filteredBands := filterBands(bands, minFormationYear, maxFormationYear, minFirstAlbumYear, maxFirstAlbumYear, minMembers, maxMembers, locations)

	// Set response headers
	w.Header().Set("Content-Type", "application/json")

	

	// Encode and return the filtered bands
	err = json.NewEncoder(w).Encode(filteredBands)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// filterBands applies the filters to the bands data
func filterBands(bands []services.Band, minFormationYear, maxFormationYear, minFirstAlbumYear, maxFirstAlbumYear, minMembers, maxMembers int, locations []string) []services.Band {
	var filtered []services.Band

	for _, band := range bands {
		// Filter by formation year
		if minFormationYear != 0 && band.FormationYear < minFormationYear {
			continue
		}
		if maxFormationYear != 0 && band.FormationYear > maxFormationYear {
			continue
		}

		// Filter by first album year
		if minFirstAlbumYear != 0 && band.FirstAlbumYear < minFirstAlbumYear {
			continue
		}
		if maxFirstAlbumYear != 0 && band.FirstAlbumYear > maxFirstAlbumYear {
			continue
		}

		// Filter by number of members
		if minMembers != 0 && band.Members < minMembers {
			continue
		}
		if maxMembers != 0 && band.Members > maxMembers {
			continue
		}

		// Filter by concert locations
		if len(locations) > 0 {
			locationMatch := false
			for _, location := range locations {
				if contains(band.ConcertLocations, location) {
					locationMatch = true
					break
				}
			}
			if !locationMatch {
				continue
			}
		}

		// If all filters pass, add the band to the filtered list
		filtered = append(filtered, band)
	}

	fmt.Println(bands)

	return filtered
}

// contains checks if a slice contains a specific value
func contains(slice []string, value string) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

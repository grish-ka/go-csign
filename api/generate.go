package handler

import (
	"encoding/json"
	"net/http"
)

// SignConfig represents the toggled state of the safety sign panels
type SignConfig struct {
	ColorMode    bool   `json:"colorMode"`
	Width        string `json:"width"`
	ActivePanels []string `json:"activePanels"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// Handle Cross-Origin Resource Sharing (CORS) so Vercel doesn't block requests
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// For now, let's make a simple health check / echo endpoint
	if r.Method == "POST" {
		var config SignConfig
		err := json.NewDecoder(r.Body).Decode(&config)
		if err != nil {
			http.Error(w, "Invalid JSON data", http.StatusBadRequest)
			return
		}

		// This is where your future Go logic will process the sign data
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "success",
			"message": "Go backend received the sign configuration successfully!",
		})
		return
	}

	// Default fallback for GET requests
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Go-CSign Backend is active and running on Vercel!"))
}
package api

import (
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"fmt"
	"github.com/hassaanaliw/seekhlai-api/scraper"
	"time"
)

// Registers all the various JSON API endpoints with the main app router
// created in main.go
func RegisterAPIRoutes(router *mux.Router) {
	router.HandleFunc("/api/v1/word/today/", GetWordToday).Methods("GET")
}

// URL: http://localhost:PORT/api/v1/word/today
// Returns the JSON of the Word of the Day from Rekhta
func GetWordToday(w http.ResponseWriter, r *http.Request) {
	word := scraper.ScrapeTodayWord(time.Now())
	ServeJson(&w, r, word)
}


// URL: http://localhost:PORT/api/v1/word/today
// Returns the JSON of the Word of the Day from Rekhta
func GetWordByDate(w http.ResponseWriter, r *http.Request) {
	word := scraper.ScrapeTodayWord(time.Now())
	ServeJson(&w, r, word)
}

// Sets the appropriate MIME types for a JSON response and modifies the headers
// Encodes the request string as a JSON object and writes it to the response
func ServeJson(w *http.ResponseWriter, r *http.Request, jsonResponse interface{}) {
	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")
	err := json.NewEncoder(*w).Encode(jsonResponse)
	if err != nil {
		// Log error if any but don't kill the request
		fmt.Println(err)
	}
}

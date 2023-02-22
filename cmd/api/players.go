package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	data2 "github.melomii/futDraft/internal/data"
)

var players struct {
	Pagination struct {
		CountCurrent   int `json:"countCurrent"`
		CountTotal     int `json:"countTotal"`
		PageCurrent    int `json:"pageCurrent"`
		PageTotal      int `json:"pageTotal"`
		ItemsPerPage   int `json:"itemsPerPage"`
	} `json:"pagination"`
	Items []data2.Players `json:"items"`
}

var player struct{
	Player struct{
		ID        int    `json:"id,omitempty"`
		CommonName string `json:"commonName,omitempty"`
		Position  string `json:"position,omitempty"`
		League  int `json:"league,omitempty"`
		Nation    int `json:"nation,omitempty"`
		Club  int `json:"club,omitempty"`
		Rating int `json:"rating,omitempty"`
		Pace int `json:"pace,omitempty"`
		Shooting int `json:"shooting,omitempty"`
		Passing int `json:"passing,omitempty"`
		Dribbling int `json:"dribbling,omitempty"`
		Defending int `json:"defending,omitempty"`
		Physicality int `json:"physicality,omitempty"`
	} `json:"player"`
}
func (app *application) ListPlayersHandler(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	if page == "" {
		page = "1" // default to page 1 if the page query parameter is not provided
	}
	
	url := fmt.Sprintf("https://futdb.app/api/players?page=%s", page)

	headers := http.Header{
		"accept":       {"application/json"},
		"X-AUTH-TOKEN": {"c0ca8775-9e7a-471a-aa7e-1cd0ecb0fa44"},
	}

	// Create the HTTP client and request
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	// Add the headers to the request
	req.Header = headers

	// Send the request and retrieve the response
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Parse the JSON response into a slice of Player structs
	err = json.NewDecoder(resp.Body).Decode(&players)
	if err != nil {
		panic(err)
	}

	jsonPlayers, err := json.Marshal(&players)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON array in the response body
	w.Write(jsonPlayers)
}

func (app *application) GetPlayerHandler(w http.ResponseWriter, r *http.Request) {

	parts := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(parts[len(parts)-1])

	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	url := fmt.Sprintf("https://futdb.app/api/players/%d", id)

	headers := http.Header{
		"accept":       {"application/json"},
		"X-AUTH-TOKEN": {"c0ca8775-9e7a-471a-aa7e-1cd0ecb0fa44"},
	}

	// Create the HTTP client and request
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	// Add the headers to the request
	req.Header = headers

	// Send the request and retrieve the response
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Parse the JSON response into a slice of Player structs
	err = json.NewDecoder(resp.Body).Decode(&player)
	if err != nil {
		panic(err)
	}

	//club,nation,leagus to string....

	
	jsonPlayer, err := json.Marshal(&player)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON array in the response body
	w.Write(jsonPlayer)
}
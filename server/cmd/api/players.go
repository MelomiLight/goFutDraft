package main

import (
	"errors"
	"fmt"
	"net/http"

	data2 "github.melomii/futDraft/internal/data"
	"github.melomii/futDraft/internal/validator"
)

func (app *application) ListPlayersHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		CommonName string
		Position  string
		data2.Filters
	}

	v := validator.New()
	qs := r.URL.Query()

	input.CommonName = app.readString(qs, "commonName", "")
	input.Position = app.readString(qs, "position", "")

	input.Filters.Sort = app.readString(qs, "sort", "id")
	input.Filters.Page = app.readInt(qs, "page", 1, v)
	input.Filters.PageSize = app.readInt(qs, "page_size", 20, v)
	input.Filters.SortSafeList = []string{"commonName", "position", "rating", "-commonName", "-position", "-rating"}

	if data2.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	players, metadata, err := app.models.Players.GetAll(input.CommonName, input.Position, input.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"players": players, "metadata": metadata}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) GetPlayerHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	players, err := app.models.Players.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data2.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"player": players}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) CreatePlayerHandler(w http.ResponseWriter, r *http.Request){
	var input struct {
		CommonName string `json:"commonName"`
		Position  string `json:"position"`
		League  int `json:"league"`
		Nation    int `json:"nation"`
		Club  int `json:"club"`
		Rating int `json:"rating"`
		Pace int `json:"pace"`
		Shooting int `json:"shooting"`
		Passing int `json:"passing"`
		Dribbling int `json:"dribbling"`
		Defending int `json:"defending"`
		Physicality int `json:"physicality"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	players := &data2.Players{
		CommonName: input.CommonName,
		Position: input.Position,
		League: input.League,  
		Nation: input.Nation,    
		Club: input.Club,  
		Rating: input.Rating,
		Pace: input.Pace, 
		Shooting: input.Shooting, 
		Passing: input.Passing, 
		Dribbling: input.Dribbling, 
		Defending: input.Defending, 
		Physicality: input.Physicality, 
	}

	err = app.models.Players.Insert(players)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/players/%d", players.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"players": players}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) DeletePlayerHandler(w http.ResponseWriter, r *http.Request){
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.models.Players.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, data2.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "player successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) UpdatePlayerHandler(w http.ResponseWriter, r *http.Request){
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	player, err := app.models.Players.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data2.ErrInvalidRuntimeFormat):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	var input struct {
		ID        int    `json:"id"`
		CommonName   *string	`json:"commonName"`
		Position    *string	`json:"position"`
		League *int	`json:"league"`
		Nation  *int	`json:"nation"`
		Club *int	`json:"club"`
		Rating *int	`json:"rating"` 
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if input.CommonName != nil {
		player.CommonName = *input.CommonName
	}

	if input.Position != nil {
		player.Position = *input.Position
	}

	if input.League != nil {
		player.League = *input.League
	}

	if input.Nation != nil {
		player.Nation = *input.Nation
	}

	if input.Club != nil {
		player.Club = *input.Club
	}

	if input.Rating != nil {
		player.Rating = *input.Rating
	}

	v := validator.New()

	if data2.ValidatePlayer(v, player); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Players.Update(player)
	if err != nil {
		switch {
		case errors.Is(err, data2.ErrEditConflict):
			app.editConflictResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"player": player}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}


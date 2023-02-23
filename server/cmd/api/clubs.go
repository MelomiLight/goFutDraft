package main

import (
	"errors"
	"net/http"

	data2 "github.melomii/futDraft/internal/data"
	"github.melomii/futDraft/internal/validator"
)

func (app *application) ListClubsHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name string
		League  int
		data2.Filters
	}

	v := validator.New()
	qs := r.URL.Query()

	input.Name = app.readString(qs, "name", "")
	input.League = app.readInt(qs, "league",13,v)

	input.Filters.Sort = app.readString(qs, "sort", "id")
	input.Filters.Page = app.readInt(qs, "page", 1, v)
	input.Filters.PageSize = app.readInt(qs, "page_size", 20, v)
	input.Filters.SortSafeList = []string{"name", "league", "id", "-name", "-league", "-id"}

	if data2.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	clubs, metadata, err := app.models.Clubs.GetAll(input.Name, input.League, input.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"clubs": clubs, "metadata": metadata}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) GetClubHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	clubs, err := app.models.Clubs.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data2.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"Club": clubs}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
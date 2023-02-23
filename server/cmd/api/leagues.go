package main

import (
	"errors"
	"net/http"

	data2 "github.melomii/futDraft/internal/data"
	"github.melomii/futDraft/internal/validator"
)

func (app *application) ListLeaguesHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name string
		data2.Filters
	}

	v := validator.New()
	qs := r.URL.Query()

	input.Name = app.readString(qs, "name", "")

	input.Filters.Sort = app.readString(qs, "sort", "id")
	input.Filters.Page = app.readInt(qs, "page", 1, v)
	input.Filters.PageSize = app.readInt(qs, "page_size", 20, v)
	input.Filters.SortSafeList = []string{"name", "id", "-name", "-id"}

	if data2.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	leagues, metadata, err := app.models.Leagues.GetAll(input.Name,input.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"leagues": leagues, "metadata": metadata}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) GetLeagueHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	leagues, err := app.models.Leagues.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data2.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"League": leagues}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
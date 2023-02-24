package main

import (
	"database/sql"
	_ "database/sql"
	"errors"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	_ "github.melomii/futDraft/internal/data"
	data2 "github.melomii/futDraft/internal/data"
)

type PlayersModels struct {
	DB *sql.DB
}



func (app *application) futDraftHandler(w http.ResponseWriter, r *http.Request) {
	user := app.contextGetUser(r)
	if user.IsAnonymous() {
		app.authenticationRequiredResponse(w, r)
		return
	}

	position433All, err := app.models.Position433All.GetAll()
	if err != nil {
		switch {
		case errors.Is(err, data2.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	
	// players,err:=app.models.Players.Get()
	// if err != nil {
	// 	switch {
	// 	case errors.Is(err, data2.ErrRecordNotFound):
	// 		app.notFoundResponse(w, r)
	// 	default:
	// 		app.serverErrorResponse(w, r, err)
	// 	}
	// 	return
	// }

	err = app.writeJSON(w, http.StatusOK, envelope{"position433": position433All}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}


}

func (app *application) GetfutDraftChooseHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	var position string
	switch id{
	case 1:
		position="GK"
	case 2:
		position="LB"
	case 3:
		position="CB"
	case 4:
		position="CB"
	case 5:
		position="RB"
	case 6:
		position="CM"
	case 7:
		position="CM"
	case 8:
		position="CM"
	case 9:
		position="LW"
	case 10:
		position="ST"
	case 11:
		position="RW"
	
	}
	positionP1, err := app.models.Players.GetRandByPosition(position)
	positionP2, err := app.models.Players.GetRandByPosition(position)
	positionP3, err := app.models.Players.GetRandByPosition(position)

	if err != nil {
		switch {
		case errors.Is(err, data2.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"player1": positionP1,"player2": positionP2,"player3": positionP3}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) PostfutDraftChooseHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		ID int `json:"id"`
		Position string `json:"position"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	players := &data2.Position433{
		ID: input.ID, 
	}
	position:=input.Position
	switch position{
	case "gk":
		err = app.models.Position433.InsertGK(players.ID)
	case "lb":
		err = app.models.Position433.InsertLB(players.ID)
	case "cb1":
		err = app.models.Position433.InsertCB1(players.ID)
	case "cb2":
		err = app.models.Position433.InsertCB2(players.ID)
	case "rb":
		err = app.models.Position433.InsertRB(players.ID)
	case "cm1":
		err = app.models.Position433.InsertCM1(players.ID)
	case "cm2":
		err = app.models.Position433.InsertCM2(players.ID)
	case "cm3":
		err = app.models.Position433.InsertCM3(players.ID)
	case "lw":
		err = app.models.Position433.InsertLW(players.ID)
	case "st":
		err = app.models.Position433.InsertST(players.ID)
	case "rw":
		err = app.models.Position433.InsertRW(players.ID)
	}
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"Message": "Choosed successfully!"},nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}






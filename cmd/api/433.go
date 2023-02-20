package main

import (
	"encoding/json"
	"net/http"

	
)

type Positions struct {
    GK  *string `json:"gk"`
    LB  *string `json:"lb"`
    CB1 *string `json:"cb1"`
    CB2 *string `json:"cb2"`
    RB  *string `json:"rb"`
    CM1 *string `json:"cm1"`
    CM2 *string `json:"cm2"`
    CM3 *string `json:"cm3"`
    LW  *string `json:"lw"`
    ST  *string `json:"st"`
    RW  *string `json:"rw"`
}

type Response struct {
    Positions Positions `json:"positions"`
    Rating    *int      `json:"rating"`
    Chemistry *int      `json:"chemistry"`
}


func (app *application) futDraftHandler(w http.ResponseWriter, r *http.Request) {

	user := app.contextGetUser(r)
	if user.IsAnonymous() {
		app.authenticationRequiredResponse(w, r)
		return
	}

	
	response := Response{
        Positions: Positions{
            GK:  nil,
            LB:  nil,
            CB1: nil,
            CB2: nil,
            RB:  nil,
            CM1: nil,
            CM2: nil,
            CM3: nil,
            LW:  nil,
            ST:  nil,
            RW:  nil,
        },
        Rating:    nil,
        Chemistry: nil,
    }

    jsonBytes, err := json.Marshal(response)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonBytes)
}
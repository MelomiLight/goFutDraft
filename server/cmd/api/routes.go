package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (applicationn *applicationn) routes() http.Handler {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(applicationn.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(applicationn.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodPost, "/register", applicationn.registerUserHandler)
	router.HandlerFunc(http.MethodGet, "/register", applicationn.showRegistrationForm)

	router.HandlerFunc(http.MethodPost, "/login", applicationn.loginUserHandler)
	router.HandlerFunc(http.MethodGet, "/login", applicationn.showLoginForm)

	router.HandlerFunc(http.MethodGet, "/futdraft/play", applicationn.futDraftHandler)
	router.HandlerFunc(http.MethodPost, "/futdraft/play/input", applicationn.PostfutDraftChooseHandler)
	router.HandlerFunc(http.MethodGet, "/futdraft/play/:id", applicationn.GetfutDraftChooseHandler)
	

	router.HandlerFunc(http.MethodGet, "/futdraft/players", applicationn.ListPlayersHandler)
	router.HandlerFunc(http.MethodGet, "/futdraft/players/:id", applicationn.GetPlayerHandler)
	router.HandlerFunc(http.MethodPost, "/futdraft/players", applicationn.CreatePlayerHandler)
	router.HandlerFunc(http.MethodDelete, "/futdraft/players/:id", applicationn.DeletePlayerHandler)
	router.HandlerFunc(http.MethodPut, "/futdraft/players/:id", applicationn.UpdatePlayerHandler)

	router.HandlerFunc(http.MethodGet, "/futdraft/clubs", applicationn.ListClubsHandler)
	router.HandlerFunc(http.MethodGet, "/futdraft/clubs/:id", applicationn.GetClubHandler)

	router.HandlerFunc(http.MethodGet, "/futdraft/leagues", applicationn.ListLeaguesHandler)
	router.HandlerFunc(http.MethodGet, "/futdraft/leagues/:id", applicationn.GetLeagueHandler)

	router.HandlerFunc(http.MethodGet, "/futdraft/nations", applicationn.ListNationsHandler)
	router.HandlerFunc(http.MethodGet, "/futdraft/nations/:id", applicationn.GetNationHandler)


	// DELETE LATER
	router.HandlerFunc(http.MethodPost, "/dbUpload", applicationn.DbUpload)
	router.HandlerFunc(http.MethodPost, "/tables", applicationn.CreateTables)

	return applicationn.recoverPanic(applicationn
.rateLimit(applicationn
.authenticate(router)))
}

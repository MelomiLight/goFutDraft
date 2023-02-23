package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodPost, "/register", app.registerUserHandler)
	router.HandlerFunc(http.MethodGet, "/register", app.showRegistrationForm)

	router.HandlerFunc(http.MethodPost, "/login", app.loginUserHandler)
	router.HandlerFunc(http.MethodGet, "/login", app.showLoginForm)

	router.HandlerFunc(http.MethodGet, "/futdraft", app.futDraftHandler)
	router.HandlerFunc(http.MethodGet, "/futdraft/choose", app.futDraftChooseHandler)

  router.HandlerFunc(http.MethodGet,"/futdraft/players", app.ListPlayersHandler)
	router.HandlerFunc(http.MethodGet,"/futdraft/players/:id", app.GetPlayerHandler)
	router.HandlerFunc(http.MethodPost,"/futdraft/players", app.CreatePlayerHandler)
	router.HandlerFunc(http.MethodDelete,"/futdraft/players/:id", app.DeletePlayerHandler)
	router.HandlerFunc(http.MethodPut,"/futdraft/players/:id", app.UpdatePlayerHandler)

	router.HandlerFunc(http.MethodGet,"/futdraft/clubs", app.ListClubsHandler)
	router.HandlerFunc(http.MethodGet,"/futdraft/clubs/:id", app.GetClubHandler)

	router.HandlerFunc(http.MethodGet,"/futdraft/leagues", app.ListLeaguesHandler)
	router.HandlerFunc(http.MethodGet,"/futdraft/leagues/:id", app.GetLeagueHandler)
	
	router.HandlerFunc(http.MethodGet,"/futdraft/nations", app.ListNationsHandler)
	router.HandlerFunc(http.MethodGet,"/futdraft/nations/:id", app.GetNationHandler)

	return app.recoverPanic(app.rateLimit(app.authenticate(router)))
}

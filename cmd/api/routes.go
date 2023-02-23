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

  router.HandlerFunc(http.MethodGet,"/futDraft/players", app.ListPlayersHandler)
	router.HandlerFunc(http.MethodGet,"/futDraft/players/:id", app.GetPlayerHandler)
	router.HandlerFunc(http.MethodPost,"/futDraft/players", app.CreatePlayerHandler)
	router.HandlerFunc(http.MethodDelete,"/futDraft/players/:id", app.DeletePlayerHandler)
	router.HandlerFunc(http.MethodPut,"/futDraft/players/:id", app.UpdatePlayerHandler)

	router.HandlerFunc(http.MethodGet,"/futDraft/clubs", app.ListClubsHandler)
	router.HandlerFunc(http.MethodGet,"/futDraft/clubs/:id", app.GetClubHandler)

	router.HandlerFunc(http.MethodGet,"/futDraft/leagues", app.ListLeaguesHandler)
	router.HandlerFunc(http.MethodGet,"/futDraft/leagues/:id", app.GetLeagueHandler)
	
	router.HandlerFunc(http.MethodGet,"/futDraft/nations", app.ListNationsHandler)
	router.HandlerFunc(http.MethodGet,"/futDraft/nations/:id", app.GetNationHandler)

	return app.recoverPanic(app.rateLimit(app.authenticate(router)))
}

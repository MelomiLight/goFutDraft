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
	router.HandlerFunc(http.MethodPost,"/login", app.loginUserHandler)
	router.HandlerFunc(http.MethodGet,"/futDraft", app.futDraftHandler)

	return app.recoverPanic(app.rateLimit(app.authenticate(router)))
}

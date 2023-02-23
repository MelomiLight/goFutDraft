package main

import (
	"context"
	"github.melomii/futDraft/internal/data"
	"net/http"
)

type contextKey string

const userContextKey = contextKey("user")

func (app *application) contextSetUser(r *http.Request, user *data.Users) *http.Request {
	ctx := context.WithValue(r.Context(), userContextKey, user)
	return r.WithContext(ctx)
}

func (app *application) contextGetUser(r *http.Request) *data.Users {
	user, ok := r.Context().Value(userContextKey).(*data.Users)
	if !ok {
	panic("missing user value in request context")
	}
	return user
	}
	
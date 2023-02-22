package main

import (
	"errors"
	"html/template"
	"net/http"
	"time"

	data2 "github.melomii/futDraft/internal/data"
	"github.melomii/futDraft/internal/validator"
)

func (app *application) 	registerUserHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	user := &data2.Users{
		Name:      input.Name,
		Email:     input.Email,
	}

	err = user.Password.Set(input.Password)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	v := validator.New()

	if data2.ValidateUser(v, user); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Users.Insert(user)

	if err != nil {
		switch {
		case errors.Is(err, data2.ErrDuplicateEmail):
			v.AddError("email", "a user with this email address already exists")
			app.failedValidationResponse(w, r, v.Errors)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"Message": "User has been created"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
	
}


func (app *application) loginUserHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()

	data2.ValidateEmail(v, input.Email)
	data2.ValidatePasswordPlaintext(v, input.Password)
	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	user, err := app.models.Users.GetByEmail(input.Email)
	if err != nil {
		switch {
		case errors.Is(err, data2.ErrRecordNotFound):
			app.invalidCredentialsResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	match, err := user.Password.Matches(input.Password)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	if !match {
		app.invalidCredentialsResponse(w, r)
		return
	}

	token, err := app.models.Tokens.New(user.ID, 24*time.Hour, data2.ScopeAuthentication)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"authentication_token": token}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
	
}

func (app *application) showRegistrationForm(w http.ResponseWriter, r *http.Request) {
	// Define the data to be passed to the template
	data := struct {
	  Title string
	}{
	  Title: "User Registration",
	}
  
	// Define the template
	tmpl := template.Must(template.ParseFiles("./public/html/reg.html"))
  
	// Render the template
	err := tmpl.Execute(w, data)
	if err != nil {
	  http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	  return
	}
  }
  func (app *application) showLoginForm(w http.ResponseWriter, r *http.Request) {
	// Define the data to be passed to the template
	data := struct {
	  Title string
	}{
	  Title: "User Registration",
	}
  
	// Define the template
	tmpl := template.Must(template.ParseFiles("./public/html/login.html"))
  
	// Render the template
	err := tmpl.Execute(w, data)
	if err != nil {
	  http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	  return
	}
  }
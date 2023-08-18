package controllers

import (
	"fmt"
	"lenslocked.com/views"
	"net/http"
)

func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
	}
}

type Users struct {
	NewView *views.View
}

// New This is used to render the form where a user can create
// a new user account.
//
// GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	u.NewView.Render(w, nil)
}

func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "This is a fake message.")
}

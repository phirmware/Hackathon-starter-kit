package controllers

import (
	"fmt"
	"hackathon/models"
	"net/http"
)

// User defines the shape of a user
type User struct {
	us models.UserDB
}

type signUpForm struct {
	FirstName string
	LastName  string
	UserName  string
	Email     string
	Password  string
}

// NewUser returns the User struct
func NewUser(us models.UserDB) *User {
	return &User{
		us: us,
	}
}

// Register handles the /signup POST
func (u *User) Register(w http.ResponseWriter, r *http.Request) {
	form := signUpForm{}
	ParseForm(r, &form)
	fmt.Fprintf(w, "%+v", form)
}

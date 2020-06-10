package controllers

import (
	"fmt"
	"hackathon/models"
	"hackathon/views"
	"net/http"
)

// User defines the shape of a user
type User struct {
	us         models.UserDB
	LoginView  *views.Views
	SignUpView *views.Views
}

type signUpForm struct {
	FirstName string `schema:"firstname"`
	LastName  string `schema:"lastname"`
	UserName  string `schema:"username"`
	Email     string `schema:"email"`
	Password  string `schema:"password"`
}

// NewUser returns the User struct
func NewUser(us models.UserDB) *User {
	return &User{
		us:         us,
		LoginView:  views.NewView("bootstrap", "user/login"),
		SignUpView: views.NewView("bootstrap", "user/signup"),
	}
}

// Login handles the /login GET
func (u *User) Login(w http.ResponseWriter, r *http.Request) {
	u.LoginView.Render(w, nil)
}

// SignUp handles the /signup GET
func (u *User) SignUp(w http.ResponseWriter, r *http.Request) {
	u.SignUpView.Render(w, nil)
}

// Register handles the /signup POST
func (u *User) Register(w http.ResponseWriter, r *http.Request) {
	form := signUpForm{}
	vd := views.Data{}
	ParseForm(r, &form)
	user := models.User{
		FirstName: form.FirstName,
		LastName:  form.LastName,
		UserName:  form.UserName,
		Email:     form.Email,
	}
	if err := u.us.Create(&user); err != nil {
		vd.Alert = &views.Alert{
			Type:    "danger",
			Message: err.Error(),
		}
		u.SignUpView.Render(w, vd)
		return
	}
	fmt.Fprintln(w, "User succesfully created")
}

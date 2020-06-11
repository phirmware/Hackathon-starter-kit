package controllers

import (
	"fmt"
	"hackathon/models"
	"hackathon/views"
	"net/http"
)

// User defines the shape of a user
type User struct {
	us         models.UserService
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

type signInForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

// NewUser returns the User struct
func NewUser(us models.UserService) *User {
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
		Password:  form.Password,
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

// SignIn handles the signin route POST /login
func (u *User) SignIn(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement this
	form := &signInForm{}
	ParseForm(r, form)
	user := &models.User{
		Email:    form.Email,
		Password: form.Password,
	}
	user, err := u.us.Authenticate(user)
	if err != nil {
		data := views.Data{}
		data.Alert = &views.Alert{
			Type:    "danger",
			Message: err.Error(),
		}
		u.LoginView.Render(w, data)
		return
	}
	fmt.Fprintln(w, user)
}

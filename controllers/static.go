package controllers

import (
	"hackathon/views"
	"net/http"
)

// Static defines the shape of the static struct
type Static struct {
	HomeView   *views.Views
	LoginView  *views.Views
	SignUpView *views.Views
}

// NewStatic returns the static struct
func NewStatic() *Static {
	return &Static{
		HomeView:   views.NewView("bootstrap", "static/home"),
		LoginView:  views.NewView("bootstrap", "user/login"),
		SignUpView: views.NewView("bootstrap", "user/signup"),
	}
}

// Home handles the / GET
func (s *Static) Home(w http.ResponseWriter, r *http.Request) {
	s.HomeView.Render(w, nil)
}

// Login handles the /login GET
func (s *Static) Login(w http.ResponseWriter, r *http.Request) {
	s.LoginView.Render(w, nil)
}

// SignUp handles the /signup GET
func (s *Static) SignUp(w http.ResponseWriter, r *http.Request) {
	s.SignUpView.Render(w, nil)
}

package controllers

import (
	"hackathon/views"
	"net/http"
)

// Static defines the shape of the static struct
type Static struct {
	HomeView *views.Views
}

// NewStatic returns the static struct
func NewStatic() *Static {
	return &Static{
		HomeView: views.NewView("bootstrap", "static/home"),
	}
}

// Home handles the / GET
func (s *Static) Home(w http.ResponseWriter, r *http.Request) {
	s.HomeView.Render(w, r, nil)
}

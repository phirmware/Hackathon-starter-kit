package controllers

import (
	"hackathon/models"
	"hackathon/views"
	"net/http"
)

// Post defines the shape of the struct
type Post struct {
	ps      models.PostService
	NewView *views.Views
}

// NewPost returns the post struct
func NewPost(ps models.PostService) *Post {
	return &Post{
		ps:      ps,
		NewView: views.NewView("bootstrap", "post/new"),
	}
}

// PostPage responds to the GET /POST route
func (p *Post) PostPage(w http.ResponseWriter, r *http.Request) {
	p.NewView.Render(w, r, nil)
}

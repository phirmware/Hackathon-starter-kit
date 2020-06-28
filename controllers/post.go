package controllers

import (
	appcontext "hackathon/context"
	"hackathon/models"
	"hackathon/views"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Post defines the shape of the struct
type Post struct {
	ps       models.PostService
	NewView  *views.Views
	ListView *views.Views
}

type postForm struct {
	Title string `schema:"title"`
	Post  string `schema:"post"`
}

// NewPost returns the post struct
func NewPost(ps models.PostService) *Post {
	return &Post{
		ps:       ps,
		NewView:  views.NewView("bootstrap", "post/new"),
		ListView: views.NewView("bootstrap", "post/list"),
	}
}

// PostPage responds to the GET /POST route
func (p *Post) PostPage(w http.ResponseWriter, r *http.Request) {
	p.NewView.Render(w, r, nil)
}

// ListPage list all users posts
func (p *Post) ListPage(w http.ResponseWriter, r *http.Request) {
	userID := appcontext.GetUserFromContext(r).ID
	posts, err := p.ps.FindByUserID(userID)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	p.ListView.Render(w, r, posts)
}

// HandlePost handles the POST /POST
func (p *Post) HandlePost(w http.ResponseWriter, r *http.Request) {
	form := &postForm{}
	ParseForm(r, form)
	userID := appcontext.GetUserFromContext(r).ID
	post := &models.Post{
		UserID: userID,
		Title:  form.Title,
		Post:   form.Post,
	}

	if err := p.ps.Create(post); err != nil {
		http.Redirect(w, r, "/post", http.StatusFound)
		return
	}

	http.Redirect(w, r, "/list", http.StatusFound)
}

// HandleDelete deletes a post from the database
func (p *Post) HandleDelete(w http.ResponseWriter, r *http.Request) {
	pd := &views.Data{}
	vars := mux.Vars(r)
	id := vars["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return
	}
	uID := uint(idInt)
	user := appcontext.GetUserFromContext(r)
	post, err := p.ps.FindPostByID(uID)
	if err != nil {
		pd.Alert = &views.Alert{
			Type:    "danger",
			Message: err.Error(),
		}
		p.ListView.Render(w, r, pd)
		return
	}
	if post.UserID != user.ID {
		pd.Alert = &views.Alert{
			Type:    "danger",
			Message: models.ErrPostMissing.Error(),
		}
		p.ListView.Render(w, r, pd)
		return
	}
	if err := p.ps.Delete(post); err != nil {
		pd.Alert = &views.Alert{
			Type:    "danger",
			Message: err.Error(),
		}
		return
	}

	http.Redirect(w, r, "/list", http.StatusFound)
}

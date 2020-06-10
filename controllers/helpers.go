package controllers

import (
	"net/http"

	"github.com/gorilla/schema"
)

// ParseForm takers form input and maps it to the data
func ParseForm(r *http.Request, form interface{}) {
	dec := schema.NewDecoder()
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	if err := dec.Decode(form, r.PostForm); err != nil {
		panic(err)
	}
}

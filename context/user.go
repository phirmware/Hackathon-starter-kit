package appcontext

import (
	"context"
	"hackathon/models"
	"net/http"
)

type usercontext string

const (
	contextUserKey usercontext = "user"
)

// SetUserInContext sets Users Context
func SetUserInContext(r *http.Request, user *models.User) *http.Request {
	ctx := context.WithValue(r.Context(), contextUserKey, user)
	r = r.WithContext(ctx)
	return r
}

// GetUserFromContext gets user from context
func GetUserFromContext(r *http.Request) *models.User {
	user := r.Context().Value(contextUserKey)
	if u, t := user.(*models.User); t {
		return u
	}
	return nil
}

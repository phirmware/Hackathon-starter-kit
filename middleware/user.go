package middleware

import (
	"context"
	"hackathon/models"
	"net/http"
)

type contextUserType string

// BrowserCookieName is the name of the cookie for the remember token
const BrowserCookieName = "remember_token"
const contextUser contextUserType = "user"

// RequireUser defines the struct for the requireuser middleware
type RequireUser struct {
	us models.UserService
}

// NewRequireUser returns the requireUser struct
func NewRequireUser(us models.UserService) *RequireUser {
	return &RequireUser{
		us: us,
	}
}

// RequireUserMiddleWare is the middleware function for user
func (ru *RequireUser) RequireUserMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(BrowserCookieName)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		user, err := ru.us.ByRemember(cookie.Value)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		ctx := context.WithValue(r.Context(), contextUser, user)
		r = r.WithContext(ctx)
		next(w, r)
	})
}

package middleware

import (
	// "context"

	appcontext "hackathon/context"
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

// User defines the shape of a user middleware
type User struct {
	us models.UserService
}

// NewRequireUser returns the requireUser struct
func NewRequireUser(us models.UserService) *RequireUser {
	return &RequireUser{
		us: us,
	}
}

// NewUser returns the User struct
func NewUser(us models.UserService) *User {
	return &User{
		us: us,
	}
}

// RequireUserMiddleWare is the middleware function for user
func (ru *RequireUser) RequireUserMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := appcontext.GetUserFromContext(r)
		if user != nil {
			next(w, r)
			return
		}
		http.Redirect(w, r, "/login", http.StatusFound)
	})
}

// UserMiddleWare is the middleware function for user
func (u *User) UserMiddleWare(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(BrowserCookieName)
		if err != nil {
			next(w, r)
			return
		}

		user, err := u.us.ByRemember(cookie.Value)
		if err != nil {
			next(w, r)
			return
		}

		r = appcontext.SetUserInContext(r, user)
		next(w, r)
	})
}

// UserMiddleWareFn function is the user middleware function
func (u *User) UserMiddleWareFn(next http.Handler) http.HandlerFunc {
	return u.UserMiddleWare(next.ServeHTTP)
}

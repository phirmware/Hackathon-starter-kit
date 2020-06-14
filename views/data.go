package views

import (
	appcontext "hackathon/context"
	"hackathon/models"
	"net/http"
)

// Alert defines the shape of the alert object
type Alert struct {
	Type    string
	Message string
}

// Data defines the shape of the data object
type Data struct {
	Alert *Alert
	User  *models.User
	Yield interface{}
}

// SetData sets the data for th templates
func SetData(r *http.Request, data interface{}) Data {
	vd := Data{}
	user := appcontext.GetUserFromContext(r)
	if d, t := data.(Data); t {
		d.User = user
		return d
	}
	vd.Yield = data
	vd.User = user
	return vd
}

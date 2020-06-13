package views

import (
	"hackathon/models"
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
func SetData(data interface{}) {
	vd := Data{}
	if _, t := data.(Data); t {
	} else {
		vd.Yield = data
	}
}

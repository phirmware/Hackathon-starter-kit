package models

import (
	"github.com/jinzhu/gorm"
)

// Services struct defines all services
type Services struct {
	User UserDB
}

// NewServices returns the services struct
func NewServices(connectionString string) (*Services, error) {
	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	return &Services{
		User: NewUserService(db),
	}, nil
}

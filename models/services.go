package models

import (
	"os"

	"github.com/jinzhu/gorm"
)

// Services struct defines all services
type Services struct {
	db   *gorm.DB
	User UserService
}

// NewServices returns the services struct
func NewServices(connectionString string) (*Services, error) {
	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	db.LogMode(true)
	return &Services{
		User: NewUserService(db),
		db:   db,
	}, nil
}

// AutoMigrate automatically creates the table in the database
func (s *Services) AutoMigrate() error {
	return s.db.AutoMigrate(&User{}).Error
}

// Close closes connection to the database
func (s *Services) Close() error {
	return s.db.Close()
}

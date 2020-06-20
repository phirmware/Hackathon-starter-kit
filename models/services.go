package models

import (
	"os"

	"github.com/jinzhu/gorm"
)

// Services struct defines all services
type Services struct {
	db   *gorm.DB
	User UserService
	Post PostService
}

// NewServices returns the services struct
func NewServices(connectionString string) (*Services, error) {
	herokuDB := os.Getenv("DATABASE_URL")
	dbURI := herokuDB
	if dbURI == "" {
		dbURI = connectionString
	}
	db, err := gorm.Open("postgres", dbURI)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)
	return &Services{
		User: NewUserService(db),
		Post: NewPostService(db),
		db:   db,
	}, nil
}

// AutoMigrate automatically creates the table in the database
func (s *Services) AutoMigrate() error {
	return s.db.AutoMigrate(&User{}, &Post{}).Error
}

// DestroyAndCreate drops all tables and recreates
func (s *Services) DestroyAndCreate() error {
	if err := s.db.DropTableIfExists(&User{}, &Post{}).Error; err != nil {
		return err
	}
	return s.AutoMigrate()
}

// Close closes connection to the database
func (s *Services) Close() error {
	return s.db.Close()
}

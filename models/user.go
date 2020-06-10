package models

import (
	"github.com/jinzhu/gorm"
)

// User defines the shape of the user table in the database
type User struct {
	gorm.Model
	FirstName    string `gorm:"not null"`
	LastName     string `gorm:"not null"`
	FullName     string `gorm:"not null"`
	Password     string `gorm:"-"`
	PasswordHash string `gorm:"not null"`
}

// UserDB defines all methods of the user service
type UserDB interface {
	Create(user *User) error
}

var _ UserDB = &userGorm{}

type userGorm struct {
	db *gorm.DB
}

func newUserGorm(db *gorm.DB) *userGorm {
	return &userGorm{
		db: db,
	}
}

// UserService defines all methods of the user service
type UserService struct {
	UserDB
}

// NewUserService returns the userservice struct
func NewUserService(db *gorm.DB) *UserService {
	ug := newUserGorm(db)
	return &UserService{
		UserDB: ug,
	}
}

func (ug *userGorm) Create(user *User) error {
	return nil
}

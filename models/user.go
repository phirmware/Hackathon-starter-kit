package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

// User defines the shape of the user table in the database
type User struct {
	gorm.Model
	FirstName    string `gorm:"not null"`
	LastName     string `gorm:"not null"`
	UserName     string `gorm:"not null"`
	Email        string `gorm:"not null"`
	Password     string `gorm:"-"`
	PasswordHash string
}

var (
	// ErrFirstNameMissing message for missing
	ErrFirstNameMissing = errors.New("models: Please provide your Firstname")
	// ErrLastNameMissing message for missing
	ErrLastNameMissing = errors.New("models: Please provide your Lastname")
	// ErrUserNameMissing message for missing
	ErrUserNameMissing = errors.New("models: Please provide a Username")
	// ErrEmailMissing message for missing
	ErrEmailMissing = errors.New("models: Please provide your Email")
	// ErrPasswordMissing message for missing
	ErrPasswordMissing = errors.New("models: Please provide a password")
)

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

type userVal struct {
	UserDB
}

func newUserVal(ug *userGorm) *userVal {
	return &userVal{
		UserDB: ug,
	}
}

// UserService defines all methods of the user service
type UserService struct {
	UserDB
}

// NewUserService returns the userservice struct
func NewUserService(db *gorm.DB) *UserService {
	ug := newUserGorm(db)
	uv := newUserVal(ug)
	return &UserService{
		UserDB: uv,
	}
}

type uservalFn func(user *User) error

func runUserValFn(user *User, fns ...uservalFn) error {
	for _, fn := range fns {
		if err := fn(user); err != nil {
			return err
		}
	}
	return nil
}

func (uv *userVal) checkforFirstName(user *User) error {
	if user.FirstName == "" {
		return ErrFirstNameMissing
	}
	return nil
}

func (uv *userVal) checkForLastName(user *User) error {
	if user.LastName == "" {
		return ErrLastNameMissing
	}
	return nil
}

func (uv *userVal) checkForUserName(user *User) error {
	if user.UserName == "" {
		return ErrUserNameMissing
	}
	return nil
}

func (uv *userVal) checkForEmail(user *User) error {
	if user.Email == "" {
		return ErrEmailMissing
	}
	return nil
}

func (uv *userVal) checkForPassword(user *User) error {
	if user.Password == "" {
		return ErrPasswordMissing
	}
	return nil
}

func (uv *userVal) Create(user *User) error {
	if err := runUserValFn(user,
		uv.checkforFirstName,
		uv.checkForLastName,
		uv.checkForUserName,
		uv.checkForEmail,
		uv.checkForPassword,
	); err != nil {
		return err
	}
	return uv.UserDB.Create(user)
}

func (ug *userGorm) Create(user *User) error {
	return ug.db.Create(&user).Error
}

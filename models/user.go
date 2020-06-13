package models

import (
	"errors"
	"hackathon/hash"
	"hackathon/rand"

	"golang.org/x/crypto/bcrypt"

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
	Remember     string
	RememberHash string
}

const (
	userPepper = "secret-user-pepper"
	hashKey    = "secret-hash-key"
)

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
	// ErrEmailExists is returned when user tries to signup with
	ErrEmailExists = errors.New("models: A user exists with that email")
	// ErrSomethingWentWrong message for missing
	ErrSomethingWentWrong = errors.New("models: Something went wrong, please try again")
	// ErrPasswordTooShort message for short password
	ErrPasswordTooShort = errors.New("models: Password provided is too short, minimum of 8 characters")
	// ErrPasswordHashMissing message for missing password hash
	ErrPasswordHashMissing = errors.New("models: Password hash missing")
	// ErrInvalidPassword message for invalid password
	ErrInvalidPassword = errors.New("models: Password Invalid")
	// ErrUserNotFound is returned when a user is not found
	ErrUserNotFound = errors.New("models: We cant find an account with that email")
	// ErrRememberNotFound is returned when a user with a rememberhash is not found
	// TODO: Change the error name, its not descriptive enough
	ErrRememberNotFound = errors.New("models: User does not exist-(remember)")
)

// UserDB defines all methods of the user service
type UserDB interface {
	Create(user *User) error
	ByEmail(user *User) (*User, error)
	ByRemember(token string) (*User, error)
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
	hmac hash.HMAC
}

func newUserVal(ug *userGorm) *userVal {
	return &userVal{
		UserDB: ug,
		hmac:   hash.NewHMAC(hashKey),
	}
}

// UserService defines all methods of the user service
type UserService interface {
	UserDB
	Authenticate(user *User) (*User, error)
}

type userService struct {
	UserDB
}

// NewUserService returns the userservice struct
func NewUserService(db *gorm.DB) UserService {
	ug := newUserGorm(db)
	uv := newUserVal(ug)
	return &userService{
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

func (uv *userVal) checkIfEmailExist(user *User) error {
	if foundUser, _ := uv.UserDB.ByEmail(user); foundUser != nil {
		return ErrEmailExists
	}
	return nil
}

func (uv *userVal) checkForPassword(user *User) error {
	if user.Password == "" {
		return ErrPasswordMissing
	}
	return nil
}

func (uv *userVal) passwordMinLength(user *User) error {
	if len(user.Password) < 8 {
		return ErrPasswordTooShort
	}
	return nil
}

func (uv *userVal) hashPassword(user *User) error {
	b, err := bcrypt.GenerateFromPassword([]byte(user.Password+userPepper), bcrypt.DefaultCost)
	if err != nil {
		return ErrSomethingWentWrong
	}
	user.PasswordHash = string(b)
	user.Password = ""
	return nil
}

func (uv *userVal) hashPasswordRequired(user *User) error {
	if user.PasswordHash == "" {
		return ErrPasswordHashMissing
	}
	return nil
}

func (uv *userVal) generateRememberToken(user *User) error {
	token, err := rand.RememberToken()
	if err != nil {
		return ErrSomethingWentWrong
	}
	user.Remember = token
	return nil
}

func (uv *userVal) rememberTokenRequired(user *User) error {
	if user.Remember == "" {
		return ErrSomethingWentWrong
	}
	return nil
}

func (uv *userVal) hashRememberToken(user *User) error {
	hash := uv.hmac.Hash(user.Remember)
	user.RememberHash = hash
	return nil
}

func (uv *userVal) rememberHashRequired(user *User) error {
	if user.RememberHash == "" {
		return ErrSomethingWentWrong
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
		uv.passwordMinLength,
		uv.checkIfEmailExist,
		uv.hashPassword,
		uv.hashPasswordRequired,
		uv.generateRememberToken,
		uv.rememberTokenRequired,
		uv.hashRememberToken,
		uv.rememberHashRequired,
	); err != nil {
		return err
	}
	return uv.UserDB.Create(user)
}

func (ug *userGorm) Create(user *User) error {
	return ug.db.Create(&user).Error
}

func (us *userService) Authenticate(user *User) (*User, error) {
	founduser, err := us.UserDB.ByEmail(user)
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(founduser.PasswordHash), []byte(user.Password+userPepper)); err != nil {
		return nil, ErrInvalidPassword
	}
	return founduser, nil
}

func (uv *userVal) ByEmail(user *User) (*User, error) {
	if err := runUserValFn(user,
		uv.checkForEmail,
	); err != nil {
		return nil, err
	}
	return uv.UserDB.ByEmail(user)
}

func (ug *userGorm) ByEmail(user *User) (*User, error) {
	if err := ug.db.First(user, "email = ?", user.Email).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (ug *userGorm) ByRemember(token string) (*User, error) {
	user := &User{
		RememberHash: token,
	}
	if err := ug.db.First(user, "remember_hash = ?", token).Error; err != nil {
		// TODO: Change the error name, its not descriptive enough
		return nil, ErrRememberNotFound
	}
	return user, nil
}

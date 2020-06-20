package models

import (
	"errors"

	"github.com/jinzhu/gorm"
)

var (
	// ErrPostMissing is returned when post field is missing
	ErrPostMissing = errors.New("models: Provide a Post")
	// ErrTitleMissing is returned when a title is miising
	ErrTitleMissing = errors.New("models: Provide a title")
)

// Post defines the shape of the post model
type Post struct {
	gorm.Model
	UserID uint   `gorm:"not null"`
	Title  string `gorm:"not null"`
	Post   string `gorm:"not null"`
}

// PostService interface
type PostService interface {
	postDB
}

type postDB interface {
	Create(post *Post) error
	FindByUserID(id uint) (*[]Post, error)
}

type postService struct {
	postDB
}

type postVal struct {
	postDB
}

type postGorm struct {
	db *gorm.DB
}

var _ postDB = &postGorm{}
var _ PostService = &postService{}

func newPostGorm(db *gorm.DB) *postGorm {
	return &postGorm{
		db: db,
	}
}

func newPostVal(pg *postGorm) *postVal {
	return &postVal{
		postDB: pg,
	}
}

func newPostService(pv *postVal) *postService {
	return &postService{
		postDB: pv,
	}
}

// NewPostService returns the PostService interface
func NewPostService(db *gorm.DB) PostService {
	pg := newPostGorm(db)
	pv := newPostVal(pg)
	return &postService{
		postDB: pv,
	}
}

type postValFn func(post *Post) error

func runPostValFns(post *Post, fns ...postValFn) error {
	for _, fn := range fns {
		if err := fn(post); err != nil {
			return err
		}
	}
	return nil
}

func (pv *postVal) checkForTitle(post *Post) error {
	if post.Title == "" {
		return ErrTitleMissing
	}
	return nil
}

func (pv *postVal) checkForPost(post *Post) error {
	if post.Post == "" {
		return ErrPostMissing
	}
	return nil
}

func (pv *postVal) Create(post *Post) error {
	if err := runPostValFns(post, pv.checkForTitle); err != nil {
		return err
	}
	return pv.postDB.Create(post)
}

func (pg *postGorm) Create(post *Post) error {
	return pg.db.Create(post).Error
}

func (pg *postGorm) FindByUserID(id uint) (*[]Post, error) {
	posts := &[]Post{}
	if err := pg.db.Find(posts, "user_id = ?", id).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

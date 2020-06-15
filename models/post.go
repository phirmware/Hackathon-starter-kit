package models

import (
	"github.com/jinzhu/gorm"
)

// Post defines the shape of the post model
type Post struct {
	gorm.Model
	Title   string `gorm:"not null"`
	Message string `gorm:"not null"`
}

// PostService interface
type PostService interface {
	postDB
}

type postDB interface {
	Create(*Post) error
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

func (pg *postGorm) Create(*Post) error {
	return nil
}

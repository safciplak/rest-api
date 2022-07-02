package comment

import "github.com/jinzhu/gorm"

// Service - the struct for our comment service
type Service struct {
	DB *gorm.DB
}

// Comment - defines our comment structure
type Comment struct {
	gorm.Model
	Slug   string
	Body   string
	Author string
}

// CommentService - the interface for our comment service
type CommentService interface {
	GetComment(ID uint) (Comment, error)
	GetCommentsBySlug(slug string) ([]Comment, error)
	PostComment(comment Comment) (Comment, error)
	UodateComment(ID uint, newComment Comment) (Comment, error)
	DeleteComment(ID uint) error
	GetAllComments() ([]Comment, error)
}

// NewService - return a new comment service
func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}

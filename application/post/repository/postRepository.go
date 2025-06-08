package repository

import (
	"core_chat/application/post/entity"
	"core_chat/application/post/model"
)

type PostRepository interface {
	CreatePost(post *model.Post) error
	GetPostByID(postID string) (*entity.Post, error)
	GetAllPosts() ([]*entity.Post, error)
}

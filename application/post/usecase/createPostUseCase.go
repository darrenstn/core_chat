package usecase

import (
	"core_chat/application/post/model"
	"core_chat/application/post/repository"
)

type CreatePostUseCase struct {
	Repo repository.PostRepository
}

func NewCreatePostUseCase(repo repository.PostRepository) *CreatePostUseCase {
	return &CreatePostUseCase{Repo: repo}
}

func (uc *CreatePostUseCase) Execute(author, title, content string) error {
	post := &model.Post{
		Author:  author,
		Title:   title,
		Content: content,
	}
	return uc.Repo.CreatePost(post)
}

package usecase

import (
	"core_chat/application/post/entity"
	"core_chat/application/post/repository"
)

type GetAllPostsUseCase struct {
	Repo repository.PostRepository
}

func NewGetAllPostsUseCase(repo repository.PostRepository) *GetAllPostsUseCase {
	return &GetAllPostsUseCase{Repo: repo}
}

func (uc *GetAllPostsUseCase) Execute() ([]*entity.Post, error) {
	return uc.Repo.GetAllPosts()
}

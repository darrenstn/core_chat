package usecase

import (
	"core_chat/application/post/entity"
	"core_chat/application/post/repository"
)

type GetPostByIDUseCase struct {
	Repo repository.PostRepository
}

func NewGetPostByIDUseCase(repo repository.PostRepository) *GetPostByIDUseCase {
	return &GetPostByIDUseCase{Repo: repo}
}

func (uc *GetPostByIDUseCase) Execute(postID string) (*entity.Post, error) {
	return uc.Repo.GetPostByID(postID)
}

package mapper

import (
	"core_chat/application/post/entity"
	"core_chat/web/rest/dto"
)

func ToPostResponse(post *entity.Post) *dto.PostResponse {
	return &dto.PostResponse{
		ID:        post.ID,
		Author:    post.Author,
		Title:     post.Title,
		Content:   post.Content,
		CreatedAt: post.CreatedAt,
	}
}

func ToPostResponseList(posts []*entity.Post) []*dto.PostResponse {
	var result []*dto.PostResponse
	for _, p := range posts {
		result = append(result, ToPostResponse(p))
	}
	return result
}

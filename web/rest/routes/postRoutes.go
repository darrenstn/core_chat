package routes

import (
	"core_chat/application/post/usecase"
	"core_chat/web/rest"
	"core_chat/web/rest/dto"
	"core_chat/web/rest/mapper"
	"core_chat/web/util"
	"encoding/json"
	"net/http"
)

type PostHandler struct {
	CreatePostUC  *usecase.CreatePostUseCase
	GetPostByIDUC *usecase.GetPostByIDUseCase
	GetAllPostsUC *usecase.GetAllPostsUseCase
}

func NewPostHandler(createPostUC *usecase.CreatePostUseCase, getPostByIDUC *usecase.GetPostByIDUseCase, getAllPostsUC *usecase.GetAllPostsUseCase) *PostHandler {
	return &PostHandler{CreatePostUC: createPostUC, GetPostByIDUC: getPostByIDUC, GetAllPostsUC: getAllPostsUC}
}

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var req dto.CreatePostRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		rest.SendResponse(w, 400, "Invalid request body")
		return
	}

	if req.Title == "" || req.Content == "" {
		rest.SendResponse(w, 400, "Title or Content should not be empty")
		return
	}

	author, ok := util.GetIdentifier(r)
	if !ok {
		rest.SendResponse(w, 401, "Unauthorized: identifier not found")
		return
	}

	err := h.CreatePostUC.Execute(author, req.Title, req.Content)
	if err != nil {
		rest.SendResponse(w, 500, "Failed to create post")
		return
	}

	rest.SendResponse(w, 200, "Post created successfully")
}

func (h *PostHandler) GetPostByID(w http.ResponseWriter, r *http.Request) {
	postID := r.URL.Query().Get("post_id")

	if postID == "" {
		rest.SendResponse(w, 400, "Post ID is required")
		return
	}

	post, err := h.GetPostByIDUC.Execute(postID)
	if err != nil {
		rest.SendResponse(w, 404, "Post not found")
		return
	}

	response := mapper.ToPostResponse(post)

	rest.SendJSON(w, struct {
		Status  int               `json:"status"`
		Message string            `json:"message"`
		Data    *dto.PostResponse `json:"data"`
	}{
		Status:  200,
		Message: "Get post by ID success",
		Data:    response,
	})
}

func (h *PostHandler) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := h.GetAllPostsUC.Execute()
	if err != nil {
		rest.SendResponse(w, 500, "Failed to fetch posts")
		return
	}

	response := mapper.ToPostResponseList(posts)

	result := struct {
		Status  int                 `json:"status"`
		Message string              `json:"message"`
		Data    []*dto.PostResponse `json:"data"`
	}{
		Status:  200,
		Message: "Success",
		Data:    response,
	}

	rest.SendJSON(w, result)
}

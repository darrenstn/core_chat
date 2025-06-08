package post

import (
	"core_chat/application/post/entity"
	"core_chat/application/post/model"
	"core_chat/application/post/repository"
	"database/sql"
	"strconv"
	"time"
)

type postRepositoryImpl struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) repository.PostRepository {
	return &postRepositoryImpl{db: db}
}

func (r *postRepositoryImpl) CreatePost(post *model.Post) error {
	_, err := r.db.Exec(`
		INSERT INTO post (author, title, content) 
		VALUES (?, ?, ?)`,
		post.Author, post.Title, post.Content)
	return err
}

func (r *postRepositoryImpl) GetPostByID(postID string) (*entity.Post, error) {
	var (
		id        int64
		createdAt time.Time
		post      entity.Post
	)

	idInt, err := strconv.ParseInt(postID, 10, 64)
	if err != nil {
		return nil, err
	}

	err = r.db.QueryRow(`
		SELECT id, author, title, content, created_at 
		FROM post WHERE id = ?`, idInt).
		Scan(&id, &post.Author, &post.Title, &post.Content, &createdAt)

	if err != nil {
		return nil, err
	}

	post.ID = strconv.FormatInt(id, 10)
	post.CreatedAt = createdAt.Format("2006-01-02 15:04:05")

	return &post, nil
}

func (r *postRepositoryImpl) GetAllPosts() ([]*entity.Post, error) {
	rows, err := r.db.Query("SELECT id, author, title, content, created_at FROM post ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*entity.Post

	for rows.Next() {
		var (
			id        int64
			createdAt time.Time
			post      entity.Post
		)

		err := rows.Scan(&id, &post.Author, &post.Title, &post.Content, &createdAt)
		if err != nil {
			return nil, err
		}

		post.ID = strconv.FormatInt(id, 10)
		post.CreatedAt = createdAt.Format("2006-01-02 15:04:05")

		posts = append(posts, &post)
	}

	return posts, nil
}

package dto

type PostResponse struct {
	ID        string `json:"id"`
	Author    string `json:"author"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
}

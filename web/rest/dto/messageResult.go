package dto

type MessageResult struct {
	ID        string `json:"id"`
	Receiver  string `json:"receiver"`
	Sender    string `json:"sender"`
	Type      string `json:"type"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	Payload   string `json:"payload"`
	CreatedAt string `json:"created_at"`
	ReadAt    string `json:"read_at"`
}

package dto

type ChatMessageResult struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Data    MessageResult `json:"data"`
}

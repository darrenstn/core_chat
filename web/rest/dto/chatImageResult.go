package dto

type ChatImageResult struct {
	Status     int    `json:"status"`
	Message    string `json:"message"`
	ChatImgURL string `json:"chat_img_url"`
}

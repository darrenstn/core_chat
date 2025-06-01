package dto

type SendNotificationInput struct {
	Receiver      string `json:"receiver"`
	Sender        string `json:"sender"`
	Type          string `json:"type"`
	Title         string `json:"title"`
	Body          string `json:"body"`
	Payload       string `json:"payload"`
	FirebaseToken string `json:"firebase_token,omitempty"`
}

package dto

type SendMessageInput struct {
	Receiver      string
	Sender        string
	Type          string
	Title         string
	Body          string
	Payload       string
	FirebaseToken string
}

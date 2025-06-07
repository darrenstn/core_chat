package dto

type SendNotificationInput struct {
	Receiver      string
	Sender        string
	Type          string
	Title         string
	Body          string
	Payload       string
	FirebaseToken string
}

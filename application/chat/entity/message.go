package entity

type Message struct {
	ID        string
	Receiver  string
	Sender    string
	Type      string
	Title     string
	Body      string
	Payload   string
	CreatedAt string
	ReadAt    string
}

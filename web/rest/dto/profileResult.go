package dto

type ProfileData struct {
	Identifier  string `json:"identifier"`
	Name        string `json:"name"`
	Description string `json:"description"`
	PictureURL  string `json:"picture_url"`
}

type ProfileResult struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    ProfileData `json:"data"`
}

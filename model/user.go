package model

type UserData struct {
	Nama string `json:"nama" binding:"required" example:"John Doe"`
}

type LoginResponse struct {
	Nama  string `json:"nama" example:"John Doe"`
	Token string `json:"token"`
}

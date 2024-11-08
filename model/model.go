package model

type FailResponse struct {
	Msg string `json:"msg" binding:"required" example:"Pesan kesalahan"`
}

type CekAuthMiddleware struct {
	Msg      string   `json:"msg" example:"Kalau middleware auth no problem: SUKSES"`
	UserData UserData `json:"user_data"`
}

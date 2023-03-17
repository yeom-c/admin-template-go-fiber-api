package handler

import (
	"time"
)

type signInReq struct {
	AuthProvider string `json:"authProvider"`
	AuthCode     string `json:"authCode"`
	Email        string `json:"email"`
	Password     string `json:"password"`
}

type signInRes struct {
	AccessToken string  `json:"access_token"`
	UserRes     userRes `json:"user"`
}

type userRes struct {
	Id        int32     `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

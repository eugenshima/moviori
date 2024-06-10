package model

import "github.com/google/uuid"

type AuthModel struct {
	Id       uuid.UUID `json:"id"`
	Login    string    `json:"login"`
	Password string    `json:"password"`
}

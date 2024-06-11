package model

import "github.com/google/uuid"

type FullUserModel struct {
	Id       uuid.UUID `json:"id"`
	Login    string    `json:"login"`
	Password []byte    `json:"password"`
}
type UserModel struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type HashedLogin struct {
	Login    string `json:"login"`
	Password []byte `json:"password"`
}

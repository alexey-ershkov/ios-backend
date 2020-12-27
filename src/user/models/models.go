package models

//easyjson:json
type User struct {
	UserID   int    `json:"userid" pg:"userid"`
	NickName string `json:"nickname" validate:"required,min=4,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=100"`
	Photo    string `json:"photo"`
}

//easyjson:json
type SafeUser struct {
	UserID   int    `json:"userid" pg:"userid"`
	NickName string `json:"name" validate:"required,min=4,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Photo    string `json:"photo"`
}

package model

import (
	"time"
)

// User Struct (Model)
type User struct {
	UserID    int64     `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	LoggedIn  bool      `json:"loggedin"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdat"`
}

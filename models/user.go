package models

type User struct {
	Username string `json:"username"`
	IsAdmin  bool   `json:"isAdmin"`
	Password string `json:"password"`
}

var (
	Users   = map[string]*User{}
	UserSeq = 1
)

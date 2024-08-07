package models

type User struct{
	UserName     string    `json:"username"`
	Role bool    `json:"role"`
	Password      string    `json:"-"`
}
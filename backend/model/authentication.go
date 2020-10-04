package model

type Authentication struct {
	Login    string `form:"login"`
	Password string `form:"password"`
}

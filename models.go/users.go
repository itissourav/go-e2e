package models

type User struct {
	Id        int64
	Firstname string
	Lastname  string
	Email     string
}

type SignupReq struct {
	Firstname string
	Lastname  string
	Email     string
	Password  string
}

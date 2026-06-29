package models

type User struct {
	Id           int64
	Firstname    string
	Lastname     string
	Email        string
	PasswordHash string
}

type SignupReq struct {
	Firstname string
	Lastname  string
	Email     string
	Password  string
}

type LoginResponse struct {
	Id           int64
	Firstname    string
	Lastname     string
	Email        string
	JwtToken     string
	RefreshToken string
}

type LoginRequest struct {
	Email    string
	Password string
}

package domain

type User struct {
	Id uint64
	Login string
	PasswordHash string
}
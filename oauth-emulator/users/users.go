package users

import (
	"math/rand"
)

type User struct {
	PersonID    int
	DisplayName string
}

type UserProvider struct {
	Users []User
}

func NewUserProvider(u ...User) *UserProvider {
	return &UserProvider{
		Users: u,
	}
}

func (u *UserProvider) RandomUser() User {
	return u.Users[rand.Intn(len(u.Users))]
}

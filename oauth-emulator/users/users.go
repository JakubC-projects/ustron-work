package users

import (
	"math/rand"

	"github.com/google/uuid"
)

type User struct {
	Uid         uuid.UUID
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

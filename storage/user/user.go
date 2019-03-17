package userTransaction

import (
	"../../helpers/password"
	"../../models/user"
)

func Register(u user.User) error {
	var err error
	return err
}

func InitializeSession(u user.User) error {
	var err error
	return err
}

func CreateJWT(u user.User) {

}

func Login(u user.User) error {
	hashedPassword := password.GenerateHashedPassword([]byte(u.Password))
}

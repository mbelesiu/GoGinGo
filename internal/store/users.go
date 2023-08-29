package store

import (
	"errors"
	"time"
)

type User struct {
	ID             int
	Username       string `binding:"required,min=5,max=30"`
	Password       string `binding:"required,min=5,max=30"`
	HashedPassword []byte `json:"-"`
	Salt           []byte `json:"-"`
	CreatedAt      time.Time
	ModifiedAt     time.Time
}

// var Users []*User
func AddUser(user *User) error {
	_, err := db.Model(user).Returning("*").Insert()
	if err != nil {
		return err
	}
	return nil
}

func Authenticate(username, password string) (*User, error) {
	user := new(User)

	if err := db.Model(user).Where(
		"username = ?", username).Select(); err != nil {
		return nil, err
	}
	if password != user.Password {
		return nil, errors.New("Invalid Password")
	}
	return user, nil
}

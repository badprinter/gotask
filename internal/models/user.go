package models

import "fmt"

type User struct {
	ID        uint
	User_name string
	Password  string
	Email     string
}

func NewUser(user_name, password, email string) *User {
	return &User{
		ID:        0,
		User_name: user_name,
		Password:  password,
		Email:     email,
	}
}

func (u *User) String() string {
	return fmt.Sprintf("User:\n\tID - %d\n\tName - %s\n\tPassword - %s\n\tEmail - %s", u.ID, u.User_name, u.Password, u.Email)
}

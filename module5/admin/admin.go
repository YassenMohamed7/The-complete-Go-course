package admin

import "example.com/module5/user"

type Admin struct {
	Email    string
	Password string
	user.User
}

func NewAdmin(email, Password, firstName, lastName, dateOfBirth string) (*Admin, error) {
	userData, err := user.New(firstName, lastName, dateOfBirth)
	if err != nil {
		return nil, err
	}
	return &Admin{
		Email:    email,
		Password: Password,
		User:     *userData,
	}, nil
}

package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName   string
	lastName    string
	dateOfBirth string
	createdAt   time.Time
}

// Constructors are special functions that are used to create and initialize instances of a struct.
// They provide a convenient way to set up the initial state of an object when it is created.
// In Go, we can define a constructor function for a struct by convention,
// typically using the name "New" followed by the struct name. "new" can't be used outside the package
// For example, we could define a constructor for the "user" struct like this:

//	we can follow the pattern making the constructor as a NewUser
//
// like func NewUser(firstName, lastName, dateOfBirth string) (*User, error){}

func New(firstName, lastName, dateOfBirth string) (*User, error) {
	if firstName == "" || lastName == "" || dateOfBirth == "" {
		return nil, errors.New("all fields are required")
	}
	return &User{
		firstName,
		lastName,
		dateOfBirth,
		time.Now(),
	}, nil
}

// this is callign by value, which means that a copy of the struct is passed to the function.
//
//	Any changes made to the struct inside the function will not affect the original struct outside the function.

func (u User) PrintUserDetailsCallingByValue() {
	fmt.Printf("Your name is %s %s\n", u.firstName, u.lastName)
	fmt.Printf("You were born on %s\n", u.dateOfBirth)
	fmt.Printf("Your account was created at %s\n", u.createdAt.Format(time.RFC1123))
}

// this is calling by reference, which means that a pointer to the struct is passed to the function.
// Any changes made to the struct inside the function will affect the original struct outside the function.

func (u *User) PrintUserDetailsCAllingByReference() {
	fmt.Printf("Your name is %s %s\n", u.firstName, u.lastName)
	fmt.Printf("You were born on %s\n", u.dateOfBirth)
	fmt.Printf("Your account was created at %s\n", u.createdAt.Format(time.RFC1123))
	u.firstName = "Aya"
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

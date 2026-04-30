package main

import (
	"fmt"

	"example.com/module5/admin"
	"example.com/module5/user"
)

func main() {
	// Before using structs, we might have used separate variables to store user data, which can be cumbersome and error-prone. For example:
	// firstName := getUserData("Enter the firstname:")
	// lastName := getUserData("Enter the lastname: ")
	// dateOfBirth := getUserData("Enter the date of birth in format DD/MM/YYYY: ")
	// fmt.Printf("Your name is %s %s and you were born on %s\n", firstName, lastName, dateOfBirth)
	// With structs, we can group related data together in a more organized way. We can define a struct to represent a user and then create an instance of that struct to store the user's information. For example:

	// This is the longer notation for struct initialization,
	// where we specify the field names explicitly.
	// This approach is more readable and less error-prone, as it clearly indicates which value corresponds to which field.
	// userData := user{
	// 	firstName:   getUserData("Enter the firstname: "),
	// 	lastName:    getUserData("Enter the lastname: "),
	// 	dateOfBirth: getUserData("Enter the date of birth in format DD/MM/YYYY: "),
	// 	createdAt:   time.Now(),
	// }

	// This called the shorter notation for struct initialization,
	// where we can directly assign values to the fields without specifying the field names.
	// However, this approach can be error-prone if the order of the values is not correct.
	// To avoid this, we can use the longer notation where we specify the field names explicitly. For example:
	// userData := user{
	// 	"Yassin",
	// 	"Mohamed",
	// 	"17/10/2000",
	// 	time.Now(),
	// }

	// userData := user{
	// 	firstName:   getUserData("Enter the firstname: "),
	// 	lastName:    getUserData("Enter the lastname: "),
	// 	dateOfBirth: getUserData("Enter the date of birth in format DD/MM/YYYY: "),
	// 	createdAt:   time.Now(),
	// }

	// Using the constructor
	userData, _ := user.New(
		getUserData("Enter the firstname: "),
		getUserData("Enter the lastname: "),
		getUserData("Enter the date of birth in format DD/MM/YYYY: "),
	)

	userData.PrintUserDetailsCallingByValue()
	userData.PrintUserDetailsCAllingByReference()
	userData.PrintUserDetailsCallingByValue()
	userData.ClearUserName()
	userData.PrintUserDetailsCallingByValue()

	adminData, _ := admin.NewAdmin(
		getUserData("Enter the email: "),
		getUserData("Enter the password: "),
		getUserData("Enter the firstname: "),
		getUserData("Enter the lastname: "),
		getUserData("Enter the date of birth in format DD/MM/YYYY: "),
	)

	fmt.Printf("Admin email: %s\n", adminData.Email)
	fmt.Printf("Admin password: %s\n", adminData.Password)
	adminData.User.PrintUserDetailsCallingByValue() // Accessing the embedded User struct's method

}

func getUserData(prompt string) string {
	fmt.Print(prompt)
	var target string
	fmt.Scan(&target)
	return target
}

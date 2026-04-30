package user

import "fmt"

type User struct {
	ID   int
	Name string
}

type Storer interface {
	Save(user User) error
}

type SqlStorer struct {
	connectionString string
}

func NewSqlStorer(connectionString string) SqlStorer {
	return SqlStorer{connectionString: connectionString}
}

func (s SqlStorer) Save(user User) error {

	fmt.Println("Saving user to SQL database:", user)
	return nil
}

type MockStorer struct {
	saveUsers []User
}

func NewMockStorer() *MockStorer {
	return &MockStorer{saveUsers: []User{}}
}
func (s *MockStorer) Save(user User) error {
	s.saveUsers = append(s.saveUsers, user)
	fmt.Println("Saving user to mock database:", s.saveUsers)
	return nil
}

func RegisterUser(storer Storer, user User) error {
	return storer.Save(user)
}

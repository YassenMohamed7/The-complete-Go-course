package main

import (
	"fmt"
	"module7/practice/example"
	"module7/practice/recursion"
	"module7/practice/retries"
	"module7/practice/server"
	"module7/practice/user"
	"time"
)

type transformFunc func(int) int

func transform(arr []int, transformer transformFunc) {
	for key, val := range arr {
		arr[key] = transformer(val)
	}
}

func getTransformFunc(multiplier int) func(int) int {
	return func(val int) int {
		return val * multiplier
	}
}

func main() {
	// fmt.Println("Advanced functions")
	// arr := []int{1, 2, 3, 4, 5}
	// transform(arr, getTransformFunc(2))
	// fmt.Println("Doubled numbers are:", arr)

	products := []example.Product{
		{Name: "Laptop", Price: 999.99, InStock: true},
		{Name: "Smartphone", Price: 499.99, InStock: false},
		{Name: "Headphones", Price: 199.99, InStock: true},
	}
	fmt.Println(example.Filter(products, func(p example.Product) bool {
		return p.InStock == true
	}))

	action := retries.WithRetries(func() error {
		return fmt.Errorf("Connection timeout!!")
	}, 3)

	error := action()
	if error != nil {
		fmt.Println("Error:", error)
	} else {
		fmt.Println("Action succeeded!")
	}

	server := server.New("localhost:8080",
		server.WithMaxConnections(30),
		server.WithTimeout(5*time.Second))
	fmt.Printf("Server configuration: %+v\n", *server)

	sql := user.NewSqlStorer("DB URL")
	mock := user.NewMockStorer()
	user.RegisterUser(sql, user.User{ID: 1, Name: "Alice"})
	user.RegisterUser(mock, user.User{ID: 2, Name: "Bob"})

	factorailOf5 := recursion.Factorial(5)
	fmt.Printf("Factorial of 5 is: %d\n", factorailOf5)

	fmt.Println("Done")

}

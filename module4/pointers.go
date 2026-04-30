package main

import "fmt"

func main() {

	x := 10
	var p *int = &x

	fmt.Println("Value of x:", x)
	fmt.Println("Address of x:", &x)
	fmt.Println("Value of p (address of x):", p)
	fmt.Println("Value pointed to by p (value of x):", *p)
	fmt.Println(&x == p)

	a := 5
	sum := 10

	add(a, &sum) // here i pass the address of sum
	// and the future changes will be reflected in sum variable
	fmt.Println("Sum after adding a:", sum)
}

func add(a int, sum *int) {
	*sum += a
}

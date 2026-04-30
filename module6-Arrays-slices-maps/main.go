package main

import "fmt"

func arrays_Slices() {

	productNames := [4]string{"Book", "Pen", "Notebook", "Eraser"}
	prices := [4]float64{10.0, 20.0, 30.0, 40.0}
	for i := 0; i < len(productNames); i++ {
		fmt.Printf("Product: %s, and Price: %.2f\n", productNames[i], prices[i])
	}

	middleProducts := productNames[1:3]   // start included but end is excluded.
	fromIndex1ToEnd := productNames[1:]   // from index 1 to the end of the array.
	formSrartToIndex2 := productNames[:2] // from the start of the array to index 2 (exclusive).
	fmt.Println("From index 1 to the end:", fromIndex1ToEnd)
	fmt.Println("From the start to index 2:", formSrartToIndex2)
	fmt.Println("Middle Products:", middleProducts)
}

type String_floatMap map[string]float64

func maps() {

	domains := map[string]string{
		"google":   "google.com",
		"facebook": "facebook.com",
		"twitter":  "twitter.com",
	}
	fmt.Println("Google: ", domains["google"])

	domains["linkedin"] = "linkedin.com"
	fmt.Println("LinkedIn: ", domains["linkedin"])
	domains["google"] = "google.co.in"
	fmt.Println("Updated Google: ", domains["google"])
	delete(domains, "google")
	fmt.Println("After Deletion, Google: ", domains)

	String_floatMap := make(map[string]float64)
	String_floatMap["pi"] = 3.14
	String_floatMap["e"] = 2.71
	fmt.Println("String to Float Map:", String_floatMap)

	for key, value := range domains {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}

}
func main() {

	// arrays_Slices()
	// exsercise()
	maps()
}

func exsercise() {
	var arr = [3]string{"Football", "Ping Pong", "Swimming"}
	for i := 0; i < len(arr); i++ {
		fmt.Printf("Sport: %s\n", arr[i])
	}
	firstAsArray := arr[0:1] // from index 0 to index 1 (exclusive
	fmt.Println("First Element standalone:", firstAsArray)
	var lastTwoElements = arr[1:3] // from index 1 to index 3 (exclusive)
	fmt.Println("Last Two Elements:", lastTwoElements)

	firstAndSecondAsSlice := firstAsArray[0:2] // from index 0 to index 2 (exclusive)
	fmt.Println("First and Second Elements as Slice:", firstAndSecondAsSlice)

	var secondAndLastAsSlice []string = firstAndSecondAsSlice[1:3] // from index 1 to index 3 (exclusive)
	fmt.Println("Second and Last Elements as Slice:", secondAndLastAsSlice)

	var courseGoals = []string{"Learn Go", "Build Projects", "Contribute to Open Source"}
	fmt.Println("Course Goals:", courseGoals)

	courseGoals[1] = "Build Backend Services"
	fmt.Println("Updated Course Goals:", courseGoals)

	products := []Product{
		{Name: "Laptop", Price: 999.99},
		{Name: "Smartphone", Price: 499.99},
	}
	fmt.Println("Products:", products)
	products = append(products, Product{Name: "Tablet", Price: 299.99})
	fmt.Println("Updated Products:", products)
	products[0].Price = 899.17
	fmt.Println("Updated Products with Discount:", products)

	prices := []float64{10.0, 20.0, 30.0}
	anotherPrices := []float64{40.0}
	prices = append(prices, anotherPrices...)
	fmt.Println("Prices after adding another prices:", prices)

}

type Product struct {
	Name  string
	Price float64
}

package main

import (
	"fmt"
	"os"
)

func main() {

	revenue := getInput("revenue")
	expenses := getInput("expenses")
	taxRate := getInput("taxRate")

	earningBeforeTax := getEBT(revenue, expenses)
	profit := getProfit(earningBeforeTax, taxRate)
	ratio := getRatio(earningBeforeTax, profit)

	printTextFormatted(earningBeforeTax, profit, ratio)
	writeCalculationsToFile(earningBeforeTax, profit, ratio)

}

func getInput(varName string) (value float64) {
	fmt.Printf("Enter the value of %v: ", varName)
	_, err := fmt.Scan(&value)

	if err != nil || value <= 0 {
		err := fmt.Sprint(varName, " must be greater than zero. Exiting program.")
		panic(err)
	}
	return
}

func getEBT(revenue float64, expenses float64) (EBT float64) {
	EBT = revenue - expenses
	return
}

func getProfit(EBT float64, taxRate float64) (profit float64) {
	profit = EBT * (1 - taxRate/100)
	return
}

func getRatio(EBT float64, profit float64) (ratio float64) {
	ratio = EBT / profit
	return
}

func printTextFormatted(EBT float64, profit float64, ratio float64) {
	fmt.Printf("Earning Before Tax (EBT) is %.2f\n", EBT)
	fmt.Printf("Earning After Tax (Profit) is %.2f\n", profit)
	fmt.Printf("Ratio (EBT / Profit) is %.2f\n", ratio)
}

func writeCalculationsToFile(EBT float64, profit float64, ratio float64) {
	data := fmt.Sprintf("EBT: %.2f\nProfit: %.2f\nRatio: %.2f\n", EBT, profit, ratio)
	err := os.WriteFile("calculations.txt", []byte(data), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}

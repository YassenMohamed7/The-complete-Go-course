package main

import (
	"fmt"

	"example.com/module3/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "account_balance.txt"

func main() {
	fmt.Println("Welcome to Go Bank App!")
	fmt.Println("Your account holder name is:", randomdata.FullName(randomdata.Male))
	balance, err := fileops.ReadFloatFromFile(accountBalanceFile)
	if err != nil {
		fmt.Println("Error reading balance:", err)
		fmt.Println("Starting with a balance of $0.00")
		balance = 0.0
		fileops.WriteFloatToFile(accountBalanceFile, balance) // wrirte initial balance to file
		// panic("Can not contiue, sorry!") This is alternative if we want to stop execution
	}
	var choice int
	for {
		printChoices()
		choice = getChoice()
		switch choice {
		case 1:
			checkBalance()
		case 2:
			balance = depositMoney(balance)
			fileops.WriteFloatToFile(accountBalanceFile, balance)
		case 3:
			balance = withdrawMoney(balance)
			fileops.WriteFloatToFile(accountBalanceFile, balance)
		case 4:
			fmt.Println("Thank you for using Go Bank App. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}

}

func getChoice() (choice int) {
	fmt.Print("Enter your choice (1-4): ")
	fmt.Scan(&choice)
	return
}

func checkBalance() {
	balance, err := fileops.ReadFloatFromFile(accountBalanceFile)
	if err != nil {
		fmt.Println("Error reading balance:", err)
		return
	}
	fmt.Printf("Your current balance is: $%.2f\n", balance)
}

func depositMoney(currnetBalance float64) float64 {
	var depositAmount float64
	fmt.Print("Enter amount to deposit: ")
	fmt.Scan(&depositAmount)
	if depositAmount <= 0 {
		fmt.Println("Deposit amount must be positive!")
		return currnetBalance
	}
	currnetBalance += depositAmount
	fmt.Printf("Successfully deposited $%.2f your available balance is now $%.2f\n", depositAmount, currnetBalance)
	return currnetBalance
}

func withdrawMoney(currentBalance float64) float64 {
	var balance float64
	fmt.Print("Enter amount to withdraw: ")
	fmt.Scan(&balance)
	if balance > currentBalance {
		fmt.Println("Insufficient funds!")
		return currentBalance
	}
	currentBalance -= balance

	fmt.Printf("Successfully withdrew $%.2f\n", balance)
	fmt.Printf("Your available balance is now $%.2f\n", currentBalance)
	return currentBalance
}

package main

import "fmt"

func main() {
	var accountBalance = 1000.0

	fmt.Println("Welcome to the Bank App")
	fmt.Println("Please select an option")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	wantsCheckBalance := choice == 1

	if wantsCheckBalance {
		fmt.Println("Your account balance is", accountBalance)

	} else if choice == 2 {
		fmt.Print("Your deposit: ")
		var depositAmount float64
		fmt.Scan(&depositAmount)
		accountBalance = accountBalance + depositAmount
		fmt.Println("Balance updated! New amount: ", accountBalance)

	} else if choice == 3 {
		fmt.Print("Withdrawal amount: ")
		var withdrawalAmount float64
		fmt.Scan(&withdrawalAmount)
		accountBalance = accountBalance - withdrawalAmount
		fmt.Println("Balance updated! New amount: ", accountBalance)
	}

	fmt.Println("Your choice is", choice)

}

package main

import (
	"fmt"

	"example.com/bank/fileops"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("----------")
		//panic("Can't countinue!")
	}

	fmt.Println("Welcome to the Bank App")

	for {
		presentOptions()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		//wantsCheckBalance := choice == 1

		switch choice {
		case 1:
			fmt.Println("Your account balance is", accountBalance)
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid deposit amount. Must be greater than 0.")
				continue
			}

			accountBalance = accountBalance + depositAmount
			fmt.Println("Balance updated! New amount: ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		case 3:
			fmt.Print("Withdrawal amount: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Invalid withdrawal amount. Must be greater than 0.")
				continue
			}

			if withdrawalAmount > accountBalance {
				fmt.Println("Insufficient funds!")
				continue
			}

			accountBalance = accountBalance - withdrawalAmount
			fmt.Println("Balance updated! New amount: ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thank you for using the Bank App")
			return
		}
	}
}

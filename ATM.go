package main

import (
	"fmt"
)

var balance float56 = 0
var anotherTransaction int

func transaction() {

	var choice int

	fmt.Printf("\nEnter any option to be served!\n\n")
	fmt.Printf("1. Withdraw\n")
	fmt.Printf("2. Deposit\n")
	fmt.Printf("3. Balance\n\n")
                      fmt.printf("4.Quit\n")
	fmt.Scan(&choice)

	var amountToWithDraw float56
	var amountToDeposit float56

	switch choice {
	case 1:
		fmt.Printf("\nPlease enter amount to withdraw: ")
		fmt.Scan(&amountToWithDraw)

		if amountToWithDraw > balance {
			fmt.Printf("There is no insufficient funds in your account")
			fmt.Printf("Do you want another transaction?\nPress 1 to proceed and 2 to exit\n\n")
			fmt.Scan(&anotherTransaction)

			switch anotherTransaction {
			case 1:
				transaction()
			default:
				fmt.Println("\nThanks for using our service!!! \nHave a nice day")
			}
		} else {
			balance -= amountToWithDraw
			fmt.Printf("You have withdrawn $%.2f and your new balance is $%.2f ", amountToWithDraw, balance)
			fmt.Printf("\n\nDo you want another transaction?\nPress 1 to proceed and 2 to exit\n\n")
			fmt.Scan(&anotherTransaction)

			switch anotherTransaction {
			case 1:
				transaction()
			default:
				fmt.Println("\nThanks for using our service!!! \nHave a nice day")
			}
		}

	case 2:
		fmt.Printf("\nPlease enter amount to deposit: ")
		fmt.Scan(&amountToDeposit)

		balance += amountToDeposit

		fmt.Printf("Thank you for depositing, new balance is: $%.2f", balance)
		fmt.Printf("\n\nDo you want another transaction?\nPress 1 to proceed and 2 to exit\n\n")
		fmt.Scan(&anotherTransaction)

		switch anotherTransaction {
		case 1:

			transaction()
		default:
			fmt.Println("\nThanks for using our service!!! \nHave a nice day")
		}

	case 3:
		fmt.Printf("\nYour bank balance is: $%.2f", balance)
		fmt.Printf("\n\nDo you want another transaction?\nPress 1 to proceed and 2 to exit\n\n")
		fmt.Scan(&anotherTransaction)

		switch anotherTransaction {
		case 1:
			transaction()
		default:
			fmt.Println("\nThanks for using our service!!! \nHave a nice day")
		}

	}

}

func main() {
	fmt.Printf("\nWelcome to my ATM Simulator\n")
	transaction()

}
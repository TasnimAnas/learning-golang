package main

import (
	"bankTerminal/fileSystem"
	"fmt"
	"time"
)

const balanceFileName = "balance.txt"
const statementFileName = "statement.csv"
const depositAction, withdrawAction = "DEPOSIT", "WITHDRAW"

func addToStatement(actionType string, actionAmount float64, newAmount float64) {
	isFileExist := fileSystem.IsFileExist(statementFileName)
	if !isFileExist {
		rowString := "Date Time, Action, Amount, Current Balance\n"
		fileSystem.ConcatStringToFile(statementFileName, rowString)
	}

	now := time.Now()
	formattedTime := now.Format("2006-01-02 15:04:05")
	rowString := fmt.Sprintf("%s, %s, %.2f, %.2f\n", formattedTime, actionType, actionAmount, newAmount)
	fileSystem.ConcatStringToFile(statementFileName, rowString)
}

func main() {
	fmt.Println("Welcome to Go Bank!")
	currentBalance := fileSystem.ReadFloatFromFile(balanceFileName)

	for {
		var option float64
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")
		fmt.Print("Choose your action: ")
		fmt.Scan(&option)

		if option == 1 {
			fmt.Println("Your balance is:", currentBalance)
			continue
		} else if option == 2 {
			var depositAmount float64
			fmt.Print("Enter deposit amount: ")
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Deposit should be greater than zero")
				continue
			}
			currentBalance += depositAmount
			fileSystem.WriteFloatToFile(balanceFileName, currentBalance)
			addToStatement(depositAction, depositAmount, currentBalance)
			fmt.Printf("Updated balance is: %.2f\n", currentBalance)
		} else if option == 3 {
			var withdrawAmount float64
			fmt.Print("Enter withdraw amount: ")
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0 {
				fmt.Println("Withdraw should be greater than zero")
				continue
			} else if withdrawAmount > currentBalance {
				fmt.Printf("Withdraw amount should be less than current balance (%.2f)\n", currentBalance)
				continue
			}
			currentBalance -= withdrawAmount
			fileSystem.WriteFloatToFile(balanceFileName, currentBalance)
			addToStatement(withdrawAction, withdrawAmount, currentBalance)
			fmt.Printf("Updated balance is: %.2f\n", currentBalance)
		} else {
			fmt.Println("Goodbye!")
			return
		}
		fmt.Printf("\n--------------------\n\n")
	}
}

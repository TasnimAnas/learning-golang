package main

import (
	"fmt"
)

func main() {
	investment := getUserInput("Enter your investment")
	revenue := getUserInput("Enter your revenue")
	tax := getUserInput("Enter your tax percentage")

	profit, profitAfterTax := calculateProfit(investment, revenue, tax)

	fmt.Printf("Your profit is: %.2f\n", profit)
	fmt.Printf("Your profit after tax is: %.2f\n", profitAfterTax)

}

func calculateProfit(investment, revenue, tax float64) (profit float64, profitAfterTax float64) {
	profit = revenue - investment
	profitAfterTax = revenue - investment - (revenue * (tax / 100))
	return profit, profitAfterTax
}

func getUserInput(showText string) (inputVal float64) {
	fmt.Printf("%s: ", showText)
	fmt.Scan(&inputVal)
	return inputVal
}

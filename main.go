// Main Go program
package main

import "fmt"

// Create a simple token definition
type cryptoCoin struct {
	coinName, tokenType string
	numberOfCoins       int
}

// Helper method to calculate token balance.
func calculateRemainingTokens(tokens int, totalTokens int) int {
	return totalTokens - tokens
}

func main() {
	// Create Fake Coin
	fakeCoin := cryptoCoin{
		coinName:      "Fako",
		tokenType:     "ERC20",
		numberOfCoins: 1000000}
	// Create Security Token
	securityToken := cryptoCoin{
		coinName:      "Asset Backed Token",
		tokenType:     "EOS",
		numberOfCoins: 25000000}
	// Print some of the field names.
	bal := calculateRemainingTokens(100000, securityToken.numberOfCoins)
	fmt.Println(securityToken.coinName, " ", bal)
	bal = calculateRemainingTokens(5000, fakeCoin.numberOfCoins)
	fmt.Println(fakeCoin.coinName, " ", bal)
}

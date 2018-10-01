// Main Go program
package main

import "fmt"

// Create a simple token definition
type cryptoCoin struct{
	coinName, tokenType  string
	numberOfCoins int64
}

func main() {
	fmt.Printf("Hiya")
	// Loop and print counter
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// Create Fake Coin
	fakeCoin := cryptoCoin{
		coinName = "Fako"
		tokenType = "ERC20"
		numberOfCoins = 1000000
	}
	// Create Security Token
	securityToken := cryptoCoin{
		coinName = "Asset Backed Token"
		tokenType = "EOS"
		numberOfCoins = 25000000
	}

}

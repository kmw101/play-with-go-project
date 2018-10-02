// Main Test
package main

import "testing"
import "strconv"

func TestMain(t *testing.T) {
	bal := calculateRemainingTokens(100, 1000)
	if bal != 900 {
		t.Errorf("The calculation of total - tokens should be 900 - result is " + strconv.Itoa(bal))
	}
	testToken := cryptoCoin{
		coinName:      "Test",
		tokenType:     "TEST",
		numberOfCoins: 100000}
	if testToken.coinName != "Test" {
		t.Errorf("coinName not set")
	}
	testToken.coinName = "Foo"
	newbal := calculatePercentageRemain(100, 0)
	if newbal != 0 {
		t.Errorf("newbal =" + strconv.Itoa(newbal))
	}
}

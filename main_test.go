// Main Test
package main

import "testing"
import "strconv"

func TestMain(t *testing.T) {
	bal := calculateRemainingTokens(100, 1000)
	if bal != 900 {
		t.Errorf("The calculation of total - tokens should be 900 - result is " + strconv.Itoa(bal))
	}

}

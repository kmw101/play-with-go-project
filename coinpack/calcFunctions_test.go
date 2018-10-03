// Test CalculatorFunctions
package coinpack

import (
	"fmt"
	"strconv"
	"testing"
)

func TestCalculateAverage(t *testing.T) {
	testsetnumbers := [10]int{10, 11, 13, 19, 20, 5, 7, 19, 30, 9}
	testres := CalculateAverage(testsetnumbers)
	fmt.Println("result returned " + strconv.Itoa(testres))
	if testres != 14 {
		t.Errorf("The calculateAverage function malfunctioned - expected 14" + strconv.Itoa(testres))
	}
}

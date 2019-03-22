// Package coinpack contains a number of calculation functions
package coinpack

import (
	"fmt"
	"strconv"
)

func init() {
	fmt.Println("starting calcfuncs")
}

// CalculateAverage takes the average of 10 ints in an array and returns an int
func CalculateAverage(nums [10]int) int {
	sum := 0
	count := len(nums)
	for i := 0; i < count; i++ {
		fmt.Print(strconv.Itoa(nums[i]))
		sum += nums[i]
	}
	return sum / count
}

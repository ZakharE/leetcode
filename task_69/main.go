package main

import (
	"fmt"
)

func main() {
	fmt.Println(mySqrt(8))
	fmt.Println(mySqrt(4))
}

func mySqrt(x int) int {
	left, right := 0, x/2
	median := left + (right-left)/2
	for {
		sqr := median * median

		if sqr > x {
			right = median - 1
		} else if sqr < x {
			left = median + 1
		}

		median = left + (right-left)/2
		fmt.Printf("l: %d, r: %d, m: %d\n", left, right, median)
		if left >= right {
			return right
		}
	}
}

// sqrt(x) 1/2x

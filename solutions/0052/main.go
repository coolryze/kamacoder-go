package main

import (
	"fmt"
)

func main() {
	n, bagWeight := 0, 0
	fmt.Scanln(&n, &bagWeight)

	weights := make([]int, n)
	values := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&weights[i], &values[i])
	}

	dp := make([]int, bagWeight+1)

	for i := 0; i < n; i++ {
		for j := weights[i]; j <= bagWeight; j++ {
			dp[j] = max(dp[j], dp[j-weights[i]]+values[i])
		}
		fmt.Println(dp)
	}

	fmt.Println(dp[bagWeight])
}

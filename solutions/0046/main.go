package main

import "fmt"

// 滚动数组
func main() {
	n := 0
	bagWeight := 0
	fmt.Scanln(&n, &bagWeight)

	weights := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&weights[i])
	}

	values := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&values[i])
	}

	dp := make([]int, bagWeight+1)
	dp[0] = 0

	for i := 0; i < n; i++ {
		for j := bagWeight; j >= weights[i]; j-- {
			dp[j] = max(dp[j], dp[j-weights[i]]+values[i])
		}
	}

	fmt.Printf("%d", dp[bagWeight])
}

// 二位数组
func main2() {
	n := 0
	bagWeight := 0
	fmt.Scanln(&n, &bagWeight)

	weights := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&weights[i])
	}

	values := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&values[i])
	}

	dp := make([][]int, n)
	for i, _ := range weights {
		dp[i] = make([]int, bagWeight+1)
	}
	for j := weights[0]; j <= bagWeight; j++ {
		dp[0][j] = values[0]
	}

	for i := 1; i < n; i++ {
		for j := 0; j <= bagWeight; j++ {
			if j < weights[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weights[i]]+values[i])
			}
		}
	}

	fmt.Printf("%d", dp[n-1][bagWeight])
}

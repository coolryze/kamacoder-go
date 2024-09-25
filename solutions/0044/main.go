package main

import (
	"fmt"
	"math"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	matrix := make([][]int, n)
	sum := 0
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&matrix[i][j])
			sum += matrix[i][j]
		}
	}

	res := math.MaxInt64
	count := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			count += matrix[i][j]
			if j == m-1 {
				res = min(res, int(math.Abs(float64(sum-count-count))))
			}
		}
	}
	count = 0
	for j := 0; j < m; j++ {
		for i := 0; i < n; i++ {
			count += matrix[i][j]
			if i == n-1 {
				res = min(res, int(math.Abs(float64(sum-count-count))))
			}
		}
	}

	fmt.Println(res)
}

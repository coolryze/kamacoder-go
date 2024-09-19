package main

import (
	"fmt"
	"math"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	grid := [][]int{}
	for i := 0; i < m; i++ {
		var s, t, v int
		fmt.Scan(&s, &t, &v)
		grid = append(grid, []int{s, t, v})
	}

	var src, dst, k int
	fmt.Scan(&src, &dst, &k)

	minDists := make([]int, n+1)
	for i := 0; i <= n; i++ {
		minDists[i] = math.MaxInt64
	}
	start := src
	end := dst
	minDists[start] = 0

	for i := 0; i < k+1; i++ {
		minDistsCopy := make([]int, n+1)
		copy(minDistsCopy, minDists)
		for _, edge := range grid {
			s, t, v := edge[0], edge[1], edge[2]
			if minDistsCopy[s] != math.MaxInt64 && minDists[t] > minDistsCopy[s]+v {
				minDists[t] = minDistsCopy[s] + v
			}
		}
	}

	if minDists[end] != math.MaxInt64 {
		fmt.Println(minDists[end])
	} else {
		fmt.Println("unconnected")
	}
}

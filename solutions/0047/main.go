package main

import (
	"fmt"
	"math"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	grid := make([][]int, n+1)
	for i := range grid {
		grid[i] = make([]int, n+1)
		for j := range grid[i] {
			grid[i][j] = math.MaxInt64
		}
	}

	for i := 0; i < m; i++ {
		var s, e, v int
		fmt.Scan(&s, &e, &v)
		grid[s][e] = v
	}

	visited := make([]bool, n+1)
	minDists := make([]int, n+1)
	for i := range minDists {
		minDists[i] = math.MaxInt64
	}
	start := 1
	end := n
	minDists[start] = 0

	for i := 1; i <= n; i++ {
		cur := 1
		minVal := math.MaxInt64
		for j := 1; j <= n; j++ {
			if !visited[j] && minDists[j] < minVal {
				cur = j
				minVal = minDists[j]
			}
		}
		visited[cur] = true

		for j := 1; j <= n; j++ {
			if !visited[j] && grid[cur][j] != math.MaxInt64 && minDists[cur]+grid[cur][j] < minDists[j] {
				minDists[j] = minDists[cur] + grid[cur][j]
			}
		}
	}

	if minDists[end] != math.MaxInt64 {
		fmt.Println(minDists[end])
	} else {
		fmt.Println(-1)
	}
}

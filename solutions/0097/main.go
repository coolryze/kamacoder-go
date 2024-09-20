package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	grid := make([][][]int, n+1)
	for i := range grid {
		grid[i] = make([][]int, n+1)
		for j := range grid[i] {
			grid[i][j] = make([]int, n+1)
			for k := range grid[i][j] {
				grid[i][j][k] = 10001
			}
		}
	}

	for i := 0; i < m; i++ {
		var u, v, w int
		fmt.Scan(&u, &v, &w)
		grid[u][v][0] = w
		grid[v][u][0] = w
	}

	var q int
	fmt.Scan(&q)
	targets := [][]int{}
	for i := 0; i < q; i++ {
		var start, end int
		fmt.Scan(&start, &end)
		targets = append(targets, []int{start, end})
	}

	for k := 1; k <= n; k++ {
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				grid[i][j][k] = min(grid[i][j][k-1], grid[i][k][k-1]+grid[k][j][k-1])
			}
		}
	}

	for _, target := range targets {
		start, end := target[0], target[1]
		if grid[start][end][n] != 10001 {
			fmt.Println(grid[start][end][n])
		} else {
			fmt.Println("-1")
		}
	}
}

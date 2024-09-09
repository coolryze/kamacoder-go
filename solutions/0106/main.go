package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	grid := make([][]int, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&grid[i][j])
		}
	}

	res := 0
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				for _, direction := range directions {
					nextI := i + direction[0]
					nextJ := j + direction[1]
					if nextI < 0 || nextI >= n || nextJ < 0 || nextJ >= m {
						res++
					}
					if grid[nextI][nextJ] == 0 {
						res++
					}
				}
			}
		}
	}

	fmt.Println(res)
}

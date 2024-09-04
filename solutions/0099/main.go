package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	grid := make([][]int, n)
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, m)
		visited[i] = make([]bool, m)

		for j := 0; j < m; j++ {
			fmt.Scan(&grid[i][j])
		}
	}

	res := 0
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	var bfs func(i, j int)
	bfs = func(i, j int) {
		queue := [][]int{{i, j}}
		visited[i][j] = true
		for len(queue) > 0 {
			cur := queue[0]
			curI := cur[0]
			curJ := cur[1]
			queue = queue[1:]
			for _, direction := range directions {
				nextI := curI + direction[0]
				nextJ := curJ + direction[1]
				if nextI < 0 || nextI >= n || nextJ < 0 || nextJ >= m {
					continue
				}
				if grid[nextI][nextJ] == 1 && !visited[nextI][nextJ] {
					queue = append(queue, []int{nextI, nextJ})
					visited[nextI][nextJ] = true
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 && !visited[i][j] {
				res++
				bfs(i, j)
			}
		}
	}

	fmt.Println(res)
}

func main2() {
	var n, m int
	fmt.Scan(&n, &m)

	grid := make([][]int, n)
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, m)
		visited[i] = make([]bool, m)

		for j := 0; j < m; j++ {
			fmt.Scan(&grid[i][j])
		}
	}

	res := 0
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	var dfs func(i, j int)
	dfs = func(i, j int) {
		for _, direction := range directions {
			nextI := i + direction[0]
			nextJ := j + direction[1]
			if nextI < 0 || nextI >= n || nextJ < 0 || nextJ >= m {
				continue
			}
			if grid[nextI][nextJ] == 1 && !visited[nextI][nextJ] {
				visited[nextI][nextJ] = true
				dfs(nextI, nextJ)
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 && !visited[i][j] {
				res++
				visited[i][j] = true
				dfs(i, j)
			}
		}
	}

	fmt.Println(res)
}

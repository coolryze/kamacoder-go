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

	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	var bfs func(int, int)
	bfs = func(i, j int) {
		queue := [][]int{}
		if grid[i][j] == 1 {
			grid[i][j] = 2
			queue = append(queue, []int{i, j})
		}
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
				if grid[nextI][nextJ] == 1 {
					grid[nextI][nextJ] = 2
					queue = append(queue, []int{nextI, nextJ})
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		bfs(i, 0)
		bfs(i, m-1)
	}
	for j := 0; j < m; j++ {
		bfs(0, j)
		bfs(n-1, j)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				grid[i][j] = 0
			} else if grid[i][j] == 2 {
				grid[i][j] = 1
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m-1; j++ {
			fmt.Printf("%d ", grid[i][j])
		}
		fmt.Println(grid[i][m-1])
	}
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

	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	var dfs func(int, int)
	dfs = func(i, j int) {
		if i < 0 || i >= n || j < 0 || j >= m || grid[i][j] == 0 || grid[i][j] == 2 {
			return
		}
		grid[i][j] = 2
		for _, direction := range directions {
			nextI := i + direction[0]
			nextJ := j + direction[1]
			dfs(nextI, nextJ)
		}
	}

	for i := 0; i < n; i++ {
		dfs(i, 0)
		dfs(i, m-1)
	}
	for j := 0; j < m; j++ {
		dfs(0, j)
		dfs(n-1, j)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				grid[i][j] = 0
			} else if grid[i][j] == 2 {
				grid[i][j] = 1
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m-1; j++ {
			fmt.Printf("%d ", grid[i][j])
		}
		fmt.Println(grid[i][m-1])
	}
}

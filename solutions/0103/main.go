package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	grid := make([][]int, n)
	firstBorder := make([][]bool, n)
	secondBorder := make([][]bool, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]int, m)
		firstBorder[i] = make([]bool, m)
		secondBorder[i] = make([]bool, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&grid[i][j])
		}
	}

	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	var dfs func(visited [][]bool, i, j int)
	dfs = func(visited [][]bool, i, j int) {
		if i < 0 || i >= n || j < 0 || j >= m || visited[i][j] {
			return
		}
		visited[i][j] = true
		for _, direction := range directions {
			nextI := i + direction[0]
			nextJ := j + direction[1]
			if nextI < 0 || nextI >= n || nextJ < 0 || nextJ >= m || grid[nextI][nextJ] < grid[i][j] {
				continue
			}
			dfs(visited, nextI, nextJ)
		}
	}

	for i := 0; i < n; i++ {
		dfs(firstBorder, i, 0)
		dfs(secondBorder, i, m-1)
	}
	for j := 0; j < m; j++ {
		dfs(firstBorder, 0, j)
		dfs(secondBorder, n-1, j)
	}

	res := [][]int{}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if firstBorder[i][j] && secondBorder[i][j] {
				res = append(res, []int{i, j})
			}
		}
	}

	for i := 0; i < len(res); i++ {
		fmt.Printf("%d %d\n", res[i][0], res[i][1])
	}
}

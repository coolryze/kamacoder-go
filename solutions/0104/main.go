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
	size := 0
	mark := 1
	markSizes := make(map[int]int)
	isAllGrid := true

	var dfs func(mark, i, j int)
	dfs = func(mark, i, j int) {
		if i < 0 || i >= n || j < 0 || j >= m || visited[i][j] || grid[i][j] != 1 {
			return
		}
		size++
		grid[i][j] = mark
		visited[i][j] = true
		for _, direction := range directions {
			nextI := i + direction[0]
			nextJ := j + direction[1]
			if nextI < 0 || nextI >= n || nextJ < 0 || nextJ >= m {
				continue
			}
			dfs(mark, nextI, nextJ)
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 0 {
				isAllGrid = false
			}
			if grid[i][j] == 1 && !visited[i][j] {
				size = 0
				mark++
				dfs(mark, i, j)
				markSizes[mark] = size
			}
		}
	}

	if isAllGrid {
		fmt.Println(m * n)
		return
	}

	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 0 {
				availableSize := 1
				marks := make(map[int]struct{})
				for _, direction := range directions {
					nextI := i + direction[0]
					nextJ := j + direction[1]
					if nextI < 0 || nextI >= n || nextJ < 0 || nextJ >= m {
						continue
					}
					nextMark := grid[nextI][nextJ]
					if _, ok := marks[nextMark]; ok {
						continue
					}
					if markSize, ok := markSizes[nextMark]; ok {
						availableSize += markSize
						marks[nextMark] = struct{}{}
					}
				}
				res = max(res, availableSize)
			}
		}
	}

	fmt.Println(res)
}

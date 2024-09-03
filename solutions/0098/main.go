package main

import "fmt"

func main() {
	var n, m int
	fmt.Scanln(&n, &m)

	graph := make([][]int, n+1)

	for i := 0; i < m; i++ {
		var s, t int
		fmt.Scanln(&s, &t)
		graph[s] = append(graph[s], t)
	}

	res := [][]int{}
	path := []int{1}

	var dfs func(graph [][]int, x int, n int)
	dfs = func(graph [][]int, x int, n int) {
		if x == n {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		for _, i := range graph[x] {
			path = append(path, i)
			dfs(graph, i, n)
			path = path[:len(path)-1]
		}
	}

	dfs(graph, 1, n)

	if len(res) == 0 {
		fmt.Println(-1)
	} else {
		for _, path := range res {
			for i := 0; i < len(path)-1; i++ {
				fmt.Printf("%d ", path[i])
			}
			fmt.Println(path[len(path)-1])
		}
	}
}

// 临接矩阵
func main1() {
	var n, m int
	fmt.Scanln(&n, &m)

	graph := make([][]int, n+1)
	for i := range graph {
		graph[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		var s, t int
		fmt.Scanln(&s, &t)
		graph[s][t] = 1
	}

	res := [][]int{}
	path := []int{1}

	var dfs func(graph [][]int, x int, n int)
	dfs = func(graph [][]int, x int, n int) {
		if x == n {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		for i := 1; i <= n; i++ {
			if graph[x][i] == 1 {
				path = append(path, i)
				dfs(graph, i, n)
				path = path[:len(path)-1]
			}
		}
	}

	dfs(graph, 1, n)

	if len(res) == 0 {
		fmt.Println(-1)
	} else {
		for _, path := range res {
			for i := 0; i < len(path)-1; i++ {
				fmt.Printf("%d ", path[i])
			}
			fmt.Println(path[len(path)-1])
		}
	}
}

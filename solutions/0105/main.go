package main

import "fmt"

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	graph := make([][]int, n+1)
	for i := 0; i < k; i++ {
		var s, t int
		fmt.Scan(&s, &t)
		graph[s] = append(graph[s], t)
	}

	visited := make([]bool, n+1)

	var dfs func(i int)
	dfs = func(i int) {
		if visited[i] {
			return
		}
		visited[i] = true
		for _, j := range graph[i] {
			dfs(j)
		}
	}

	dfs(1)

	res := 1
	for i := 1; i <= n; i++ {
		if !visited[i] {
			res = -1
		}
	}

	fmt.Println(res)
}

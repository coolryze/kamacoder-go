package main

import (
	"fmt"
	"math"
)

// Bellman-ford 队列优化
func main() {
	var n, m int
	fmt.Scan(&n, &m)

	grid := make([][][]int, n+1)
	for i := 0; i < m; i++ {
		var s, t, v int
		fmt.Scan(&s, &t, &v)
		grid[s] = append(grid[s], []int{t, v})
	}

	start := 1
	end := n
	queue := []int{1}
	minDists := make([]int, n+1)
	for i := 0; i <= n; i++ {
		minDists[i] = math.MaxInt64
	}
	minDists[start] = 0
	counter := make([]int, n+1)
	flag := false

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		for _, edge := range grid[node] {
			from := node
			to := edge[0]
			val := edge[1]
			if minDists[to] > minDists[from]+val {
				minDists[to] = minDists[from] + val
				queue = append(queue, to)
				counter[to]++
				if counter[to] == n {
					flag = true
					for len(queue) > 0 {
						queue = []int{}
					}
					break
				}
			}
		}
	}

	if flag {
		fmt.Println("circle")
	} else if minDists[end] == math.MaxInt64 {
		fmt.Println("unconnected")
	} else {
		fmt.Println(minDists[end])
	}
}

// Bellman-ford
func main2() {
	var n, m int
	fmt.Scan(&n, &m)

	grid := [][]int{}
	for i := 0; i < m; i++ {
		var s, t, v int
		fmt.Scan(&s, &t, &v)
		grid = append(grid, []int{s, t, v})
	}

	minDists := make([]int, n+1)
	for i := 0; i <= n; i++ {
		minDists[i] = math.MaxInt64
	}
	start := 1
	end := n
	minDists[start] = 0
	flag := false

	for i := 0; i <= n; i++ {
		updated := false
		for _, edge := range grid {
			s, t, v := edge[0], edge[1], edge[2]
			if minDists[s] != math.MaxInt64 && minDists[s]+v < minDists[t] {
				if i < n {
					minDists[t] = minDists[s] + v
					updated = true
				} else {
					flag = true
				}
			}
		}
		if !updated {
			break
		}
	}

	if flag {
		fmt.Println("circle")
	} else if minDists[end] == math.MaxInt64 {
		fmt.Println("unconnected")
	} else {
		fmt.Println(minDists[end])
	}
}

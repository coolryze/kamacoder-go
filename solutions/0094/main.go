package main

import (
	"fmt"
	"math"
)

func main() {
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

	for i := 1; i <= n; i++ {
		updated := false
		for _, edge := range grid {
			s, t, v := edge[0], edge[1], edge[2]
			if minDists[s] != math.MaxInt64 && minDists[s]+v < minDists[t] {
				minDists[t] = minDists[s] + v
				updated = true
			}
		}
		if !updated {
			break
		}
	}

	if minDists[end] != math.MaxInt64 {
		fmt.Println(minDists[end])
	} else {
		fmt.Println("unconnected")
	}
}

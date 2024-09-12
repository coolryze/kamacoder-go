package main

import "fmt"

const maxNode = 1000

var n int
var father [maxNode + 1]int

func initialize() {
	for i := 1; i <= n; i++ {
		father[i] = i
	}
}

func find(u int) int {
	if u == father[u] {
		return u
	}
	father[u] = find(father[u])
	return father[u]
}

func isSame(u, v int) bool {
	u = find(u)
	v = find(v)
	return u == v
}

func join(u, v int) {
	u = find(u)
	v = find(v)
	if u == v {
		return
	}
	father[v] = u
}

func isTreeAfterRemoveEdge(edges [][]int, deleteEdge int) bool {
	initialize()
	for i := 0; i < n; i++ {
		if i == deleteEdge {
			continue
		}
		if isSame(edges[i][0], edges[i][1]) {
			return false
		} else {
			join(edges[i][0], edges[i][1])
		}
	}
	return true
}

func getRemoveEdge(edges [][]int) (int, int) {
	initialize()
	for i := 0; i < n; i++ {
		if isSame(edges[i][0], edges[i][1]) {
			return edges[i][0], edges[i][1]
		} else {
			join(edges[i][0], edges[i][1])
		}
	}
	return -1, -1
}

func main() {
	fmt.Scan(&n)
	edges := make([][]int, n)
	inDegree := make([]int, n+1)
	vec := []int{}

	for i := 0; i < n; i++ {
		var s, t int
		fmt.Scan(&s, &t)
		edges[i] = []int{s, t}
		inDegree[t]++
	}

	for i := n - 1; i >= 0; i-- {
		if inDegree[edges[i][1]] == 2 {
			vec = append(vec, i)
		}
	}

	if len(vec) > 0 {
		if isTreeAfterRemoveEdge(edges, vec[0]) {
			fmt.Printf("%d %d\n", edges[vec[0]][0], edges[vec[0]][1])
		} else {
			fmt.Printf("%d %d\n", edges[vec[1]][0], edges[vec[1]][1])
		}
	} else {
		s, t := getRemoveEdge(edges)
		fmt.Printf("%d %d\n", s, t)
	}
}

package main

import (
	"fmt"
	"math"
	"sort"
)

const maxNode = 10000

var v int
var father = [maxNode + 1]int{}

type edge struct {
	v1  int
	v2  int
	val int
}

func initialize() {
	for i := 1; i <= v; i++ {
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

func join(u, v int) {
	u = find(u)
	v = find(v)
	if u == v {
		return
	}
	father[v] = u
}

// kruskal
func main() {
	var e int
	fmt.Scan(&v, &e)

	edges := []edge{}
	for i := 0; i < e; i++ {
		var v1, v2, val int
		fmt.Scan(&v1, &v2, &val)
		edges = append(edges, edge{v1, v2, val})
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].val < edges[j].val
	})

	initialize()
	res := 0
	for i := 0; i < e; i++ {
		edge := edges[i]
		if find(edge.v1) != find(edge.v2) {
			join(edge.v1, edge.v2)
			res += edge.val
		}
	}

	fmt.Println(res)
}

// prim
func main2() {
	var v, e int
	fmt.Scan(&v, &e)

	defaultVal := 10000

	grid := make([][]int, v+1)
	for i := range grid {
		grid[i] = make([]int, v+1)
		for j := range grid[i] {
			grid[i][j] = defaultVal
		}
	}
	for i := 0; i < e; i++ {
		var v1, v2, val int
		fmt.Scan(&v1, &v2, &val)
		grid[v1][v2] = val
		grid[v2][v1] = val
	}

	inTrees := make([]bool, v+1)
	minDists := make([]int, v+1)
	for i := 0; i <= v; i++ {
		minDists[i] = defaultVal
	}

	for i := 1; i < v; i++ {
		cur := -1
		minVal := math.MaxInt64
		for j := 1; j <= v; j++ {
			if !inTrees[j] && minDists[j] < minVal {
				cur = j
				minVal = minDists[j]
			}
		}
		inTrees[cur] = true

		for j := 1; j <= v; j++ {
			if !inTrees[j] && grid[cur][j] < minDists[j] {
				minDists[j] = grid[cur][j]
			}
		}
	}

	res := 0
	for i := 2; i <= v; i++ {
		res += minDists[i]
	}

	fmt.Println(res)
}

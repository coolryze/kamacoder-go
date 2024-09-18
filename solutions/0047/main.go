package main

import (
	"container/heap"
	"fmt"
	"math"
)

type edge struct {
	to  int
	val int
}

type entry struct {
	node int
	dist int
}

type minHeap []entry

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i].dist < h[j].dist }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(entry)) }
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// dijkstra 堆排序版
func main() {
	var n, m int
	fmt.Scan(&n, &m)

	grid := make([][]edge, n+1)
	for i := 0; i < m; i++ {
		var s, e, v int
		fmt.Scan(&s, &e, &v)
		grid[s] = append(grid[s], edge{e, v})
	}

	visited := make([]bool, n+1)
	minDists := make([]int, n+1)
	for i := range minDists {
		minDists[i] = math.MaxInt64
	}
	start := 1
	end := n

	h := &minHeap{}
	heap.Init(h)
	heap.Push(h, entry{start, 0})
	minDists[start] = 0

	for h.Len() > 0 {
		cur := heap.Pop(h).(entry)
		if visited[cur.node] {
			continue
		}
		visited[cur.node] = true

		for _, e := range grid[cur.node] {
			if !visited[e.to] && minDists[cur.node]+e.val < minDists[e.to] {
				minDists[e.to] = minDists[cur.node] + e.val
				heap.Push(h, entry{e.to, minDists[e.to]})
			}
		}
	}

	if minDists[end] != math.MaxInt64 {
		fmt.Println(minDists[end])
	} else {
		fmt.Println(-1)
	}
}

// dijkstra 朴素版
func main2() {
	var n, m int
	fmt.Scan(&n, &m)

	grid := make([][]int, n+1)
	for i := range grid {
		grid[i] = make([]int, n+1)
		for j := range grid[i] {
			grid[i][j] = math.MaxInt64
		}
	}

	for i := 0; i < m; i++ {
		var s, e, v int
		fmt.Scan(&s, &e, &v)
		grid[s][e] = v
	}

	visited := make([]bool, n+1)
	minDists := make([]int, n+1)
	for i := range minDists {
		minDists[i] = math.MaxInt64
	}
	start := 1
	end := n
	minDists[start] = 0

	for i := 1; i <= n; i++ {
		cur := 1
		minVal := math.MaxInt64
		for j := 1; j <= n; j++ {
			if !visited[j] && minDists[j] < minVal {
				cur = j
				minVal = minDists[j]
			}
		}
		visited[cur] = true

		for j := 1; j <= n; j++ {
			if !visited[j] && grid[cur][j] != math.MaxInt64 && minDists[cur]+grid[cur][j] < minDists[j] {
				minDists[j] = minDists[cur] + grid[cur][j]
			}
		}
	}

	if minDists[end] != math.MaxInt64 {
		fmt.Println(minDists[end])
	} else {
		fmt.Println(-1)
	}
}

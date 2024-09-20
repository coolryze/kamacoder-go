package main

import (
	"container/heap"
	"fmt"
	"math"
)

// 定义一对坐标 (x, y)
type Point struct {
	x, y int
}

// 定义一个用于优先队列的结构体
type Item struct {
	point    Point
	distance float64
	steps    int
}

// 优先队列的类型
type PriorityQueue []Item

// 实现 heap.Interface 接口的三个方法
func (pq PriorityQueue) Len() int            { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool  { return pq[i].distance < pq[j].distance }
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(Item)) }
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// 骑士的所有可能移动
var moves = []Point{
	{1, 2}, {2, 1}, {-1, 2}, {2, -1},
	{1, -2}, {-2, 1}, {-1, -2}, {-2, -1},
}

// 计算欧几里得距离
func distance(a, b Point) float64 {
	return math.Sqrt(float64((a.x-b.x)*(a.x-b.x) + (a.y-b.y)*(a.y-b.y)))
}

// 使用启发式搜索（类似A*）查找最短路径
func bfs(start, end Point) int {
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, Item{point: start, distance: distance(start, end), steps: 0})

	step := make(map[Point]int)
	step[start] = 0

	for pq.Len() > 0 {
		curItem := heap.Pop(pq).(Item)
		cur := curItem.point
		if cur == end {
			return curItem.steps
		}
		for _, move := range moves {
			newPoint := Point{cur.x + move.x, cur.y + move.y}
			if newPoint.x >= 1 && newPoint.x <= 1000 && newPoint.y >= 1 && newPoint.y <= 1000 {
				newSteps := curItem.steps + 1
				if existingSteps, exists := step[newPoint]; !exists || newSteps < existingSteps {
					step[newPoint] = newSteps
					heap.Push(pq, Item{point: newPoint, distance: distance(newPoint, end) + float64(newSteps), steps: newSteps})
				}
			}
		}
	}
	return -1 // 无法到达
}

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var a1, a2, b1, b2 int
		fmt.Scan(&a1, &a2, &b1, &b2)
		start := Point{a1, a2}
		end := Point{b1, b2}
		fmt.Println(bfs(start, end))
	}
}

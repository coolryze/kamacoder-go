package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	umap := make([][]int, n)
	inDegree := make([]int, n)
	for i := 0; i < m; i++ {
		var s, t int
		fmt.Scan(&s, &t)
		umap[s] = append(umap[s], t)
		inDegree[t]++
	}

	queue := []int{}
	for i := 0; i < n; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	res := []int{}
	for len(queue) > 0 {
		s := queue[0]
		queue = queue[1:]
		res = append(res, s)
		for _, t := range umap[s] {
			inDegree[t]--
			if inDegree[t] == 0 {
				queue = append(queue, t)
			}
		}
	}

	if len(res) == n {
		for i := 0; i < len(res)-1; i++ {
			fmt.Printf("%d ", res[i])
		}
		fmt.Println(res[len(res)-1])
	} else {
		fmt.Println(-1)
	}
}

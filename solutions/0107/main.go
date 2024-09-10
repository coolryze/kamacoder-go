package main

import "fmt"

const maxNode = 100

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

func main() {
	var m int
	fmt.Scan(&n, &m)

	initialize()

	for i := 0; i < m; i++ {
		var s, t int
		fmt.Scan(&s, &t)
		join(s, t)
	}

	var source, destination int
	fmt.Scan(&source, &destination)

	if isSame(source, destination) {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}

package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	nums := make([]int, n)
	preSum := 0
	preSums := make([]int, n)
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		nums[i] = num
		preSum += num
		preSums[i] = preSum
	}

	indexes := [][]int{}
	for {
		var i, j int
		_, err := fmt.Scanf("%d %d", &i, &j)
		if err != nil {
			break
		}
		indexes = append(indexes, []int{i, j})
	}

	for _, index := range indexes {
		i, j := index[0], index[1]
		if i == 0 {
			fmt.Println(preSums[j])
		} else {
			fmt.Println(preSums[j] - preSums[i-1])
		}
	}
}

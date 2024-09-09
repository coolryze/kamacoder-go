package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var beginStr, endStr string
	fmt.Scan(&beginStr, &endStr)

	strList := make(map[string]struct{}, n)
	for i := 0; i < n; i++ {
		var str string
		fmt.Scan(&str)
		strList[str] = struct{}{}
	}

	queue := []string{beginStr}
	visited := map[string]int{beginStr: 1}

	res := 0

	for len(queue) > 0 {
		curStr := queue[0]
		queue = queue[1:]
		path := visited[curStr]

		for i := 0; i < len(curStr); i++ {
			chars := []rune(curStr)
			for j := 0; j < 26; j++ {
				chars[i] = 'a' + rune(j)
				newStr := string(chars)
				if newStr == endStr {
					res = path + 1
					break
				}
				if _, ok := strList[newStr]; ok {
					if _, ok := visited[newStr]; !ok {
						queue = append(queue, newStr)
						visited[newStr] = path + 1
					}
				}
			}
			if res > 0 {
				break
			}
		}
	}

	fmt.Println(res)
}

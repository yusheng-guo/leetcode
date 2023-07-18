package main

import "fmt"

func transportationHub(path [][]int) int {
	l := len(path)
	out := map[int]struct{}{}      // 出
	in := map[int]int{}            // 入
	location := map[int]struct{}{} // 所有地点
	for i := 0; i < l; i++ {
		if _, ok := out[path[i][0]]; !ok {
			out[path[i][0]] = struct{}{}
		}
		if _, ok := location[path[i][0]]; !ok {
			location[path[i][0]] = struct{}{}
		}
		if _, ok := location[path[i][1]]; !ok {
			location[path[i][1]] = struct{}{}
		}
		in[path[i][1]]++
	}
	fmt.Println(out, in, location)
	for k, v := range in {
		_, ok := out[k]
		if v == len(location)-1 && !ok {
			return k
		}
	}
	return -1
}

// func main() {
// 	//path := [][]int{{0, 1}, {0, 3}, {1, 3}, {2, 0}, {2, 3}}
// 	path := [][]int{{1, 3}, {1, 2}, {0, 2}, {3, 2}}
// 	ret := transportationHub(path)
// 	fmt.Println(ret)
// }

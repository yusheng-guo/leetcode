package main

import (
	"strings"
)

// 2. 探险营地
func adventureCamp(expeditions []string) int {
	max := 0
	index := -1
	hasVisited := map[string]struct{}{}
	for i, v := range expeditions {
		if v == "" { //未发现新营地
			continue
		}
		ss := strings.Split(v, "->")
		if i == 0 { // 初始小扣已知的所有营地
			for _, s := range ss {
				hasVisited[s] = struct{}{}
			}
			continue
		}
		count := 0 // 第i次探险 新发现的营地数量
		for _, s := range ss {
			if _, ok := hasVisited[s]; !ok {
				hasVisited[s] = struct{}{}
				count++
			}
		}
		if max < count {
			max = count
			index = i
		}
	}
	return index
}

// func main() {
// 	expeditions := []string{"Alice->Dex", "", "Dex"}
// 	//expeditions := []string{"", "Gryffindor->Slytherin->Gryffindor", "Hogwarts->Hufflepuff->Ravenclaw"}
// 	fmt.Println(adventureCamp(expeditions))
// }

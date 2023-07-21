package main

import "fmt"

// 2383. 赢得比赛需要的最少训练时长
// https://leetcode.cn/problems/minimum-hours-of-training-to-win-a-competition/
func minNumberOfHours(initialEnergy int, initialExperience int, energy []int, experience []int) int {
	// ① 精力 满足 需要最少时长
	sumEnergy := 0     // 击败所有对手需要的总精力
	requireEnergy := 0 // 除了初始精力之外 还需要多少精力 = 需要训练多少年
	for i := 0; i < len(energy); i++ {
		sumEnergy += energy[i]
	}
	if initialEnergy <= sumEnergy {
		requireEnergy = sumEnergy - initialEnergy + 1
	}
	fmt.Println(requireEnergy)
	// ② 经验 满足 需要多长时间
	currentExperience := initialExperience // 当前剩余经验值
	requireExperience := 0                 // 除了初始经验 还需要多少经验 = 需要训练多少年
	for i := 0; i < len(experience); i++ {
		if currentExperience <= experience[i] { // 打不多
			requireExperience += experience[i] - currentExperience + 1
			currentExperience = +experience[i]*2 + 1 // 更新 currentExperience
		} else { // 打过了 更新 currentExperience
			currentExperience += experience[i]
		}
	}
	fmt.Println(requireExperience)
	return requireEnergy + requireExperience
}

// func main() {
// 	initialEnergy := 2
// 	initialExperience := 4
// 	energy := []int{1}
// 	experience := []int{3}
// 	ret := minNumberOfHours(initialEnergy, initialExperience, energy, experience)
// 	fmt.Println(ret)
// }

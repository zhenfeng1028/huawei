package main

import (
	"fmt"
	"sort"
)

func main() {
	var m, n int // m为核酸采样人员数量，n为志愿者数量
	fmt.Scan(&m, &n)
	sampleEfficiency := make([]int, m) // 采样员的效率值
	for i := 0; i < m; i++ {
		var efficiency int
		fmt.Scan(&efficiency)
		sampleEfficiency[i] = efficiency
	}
	sort.SliceStable(sampleEfficiency, func(i, j int) bool {
		return sampleEfficiency[i] > sampleEfficiency[j]
	})
	res := 0
	increase := []int{} // 配备志愿者后的增益
	for _, efficiency := range sampleEfficiency {
		res += efficiency
		increase = append(increase, int(0.2*float64(efficiency))) // 有一名志愿者协助组织，效率提升2M
		for i := 0; i < 3; i++ {
			increase = append(increase, int(0.1*float64(efficiency))) // 每增加一名志愿者，效率提升1M，最多提升3M
		}
	}
	sort.SliceStable(increase, func(i, j int) bool {
		return increase[i] > increase[j] // 将志愿者分配到采样员效率最大化
	})
	res = int(0.8 * float64(res)) // 如果没有志愿者协助组织，效率下降2M
	for i := 0; i < n; i++ {
		if i >= len(increase) { // 多余的志愿者无法进行分配
			break
		}
		res += increase[i] // 加上志愿者带来的增益
	}
	fmt.Println(res)
}

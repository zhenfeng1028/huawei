package main

import "fmt"

var n, k int                // 服务总数，待求的服务k
var services [][]int        // 各服务启动时长及依赖关系
var hashmap = map[int]int{} // 各服务可以进行集成测试时的时长

func main() {
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var service []int
		for j := 0; j < n; j++ {
			var tmp int
			fmt.Scan(&tmp)
			service = append(service, tmp)
		}
		services = append(services, service)
	}
	fmt.Scan(&k)
	fmt.Println(computeTime(k - 1))
}

func computeTime(k int) int {
	var startUpTime, dependentTime int // 服务自身的启动时间和该服务依赖服务启动所需时间
	for i := 0; i < n; i++ {
		if i == k {
			startUpTime = services[k][i]
			continue
		}
		if services[k][i] == 1 {
			if _, ok := hashmap[i]; !ok {
				hashmap[i] = computeTime(i)
			}
			dependentTime = max(dependentTime, hashmap[i])
		}
	}
	return startUpTime + dependentTime
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

package main

import (
	"fmt"
	"sort"
)

type Region struct {
	Left  float64
	Right float64
}

var region []*Region

func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var left, right float64
		fmt.Scan(&left, &right)
		region = append(region, &Region{Left: left, Right: right})
	}
	sort.Sort(RegionSlice(region))
	left, right := region[0].Left, region[1].Right
	res := 0.0
	for left < right {
		mid := float64(left) + float64(right-left)/2
		disMid := distance(mid)
		disLeft := distance(mid - 0.5)
		disRight := distance(mid + 0.5)
		if disLeft < disMid {
			right = mid - 0.5
			res = disLeft
		} else {
			left = mid + 0.5
			res = disRight
		}
	}
	fmt.Println(int(res))
}

func distance(location float64) float64 {
	var sum float64
	for _, r := range region {
		if location < r.Left {
			sum += r.Left - location
		} else if r.Right < location {
			sum += location - r.Right
		}
	}
	return sum
}

type RegionSlice []*Region

func (s RegionSlice) Len() int      { return len(s) }
func (s RegionSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s RegionSlice) Less(i, j int) bool {
	if s[i].Left != s[j].Left {
		return s[i].Left < s[j].Left
	} else {
		return s[i].Right < s[j].Right
	}
}

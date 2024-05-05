package main

import (
	"fmt"
	"sort"
)

type Record struct {
	Id    int `json:"id"`    // 员工id
	Count int `json:"count"` // 打卡次数
	First int `json:"first"` // 第一次打卡时间
}

var whole []int    // 30天打卡整体情况
var detail [][]int // 30天打卡详细情况

func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < 30; i++ {
		var dayNum int
		fmt.Scan(&dayNum)
		whole = append(whole, dayNum)
	}
	for i := 0; i < 30; i++ {
		var dayDetail []int
		for j := 0; j < whole[i]; j++ {
			var id int
			fmt.Scan(&id)
			dayDetail = append(dayDetail, id)
		}
		detail = append(detail, dayDetail)
	}
	hashmap := make(map[int]*Record)
	for i, dayDetail := range detail {
		for _, id := range dayDetail {
			if _, ok := hashmap[id]; !ok {
				hashmap[id] = &Record{
					Id:    id,
					Count: 1,
					First: i,
				}
			} else {
				hashmap[id].Count++
			}
		}
	}
	records := []*Record{}
	for _, v := range hashmap {
		records = append(records, v)
	}
	sort.Sort(RecordSlice(records))
	for i, v := range records {
		if i >= 5 {
			break
		}
		fmt.Printf("%d ", v.Id)
	}
}

type RecordSlice []*Record

func (s RecordSlice) Len() int      { return len(s) }
func (s RecordSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s RecordSlice) Less(i, j int) bool {
	if s[i].Count != s[j].Count {
		return s[i].Count > s[j].Count
	} else if s[i].First != s[j].First {
		return s[i].First < s[j].First
	} else {
		return s[i].Id < s[j].Id
	}
}

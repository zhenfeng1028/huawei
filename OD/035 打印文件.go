package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Event struct {
	Id       int
	Priority int
}

func main() {
	var n int
	fmt.Scan(&n)
	var events [6][]*Event
	for i := 0; i < 6; i++ {
		events[i] = make([]*Event, 0)
	}
	input := bufio.NewScanner(os.Stdin)
	res := []string{}
	for i := 0; i < n; i++ {
		event := &Event{Id: i + 1}
		var op string
		var printer, priority int
		input.Scan()
		strs := strings.Split(input.Text(), " ")
		op = strs[0]
		printer, _ = strconv.Atoi(strs[1])
		if len(strs) == 3 {
			priority, _ = strconv.Atoi(strs[2])
		}
		event.Priority = priority
		if op == "IN" {
			events[printer] = append(events[printer], event)
			sort.Sort(EventSlice(events[printer]))
		} else {
			if len(events[printer]) == 0 {
				res = append(res, "NULL")
			} else {
				res = append(res, strconv.Itoa(events[printer][0].Id))
				events[printer] = events[printer][1:]
			}
		}
	}
	for _, v := range res {
		fmt.Println(v)
	}
}

type EventSlice []*Event

func (s EventSlice) Len() int      { return len(s) }
func (s EventSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s EventSlice) Less(i, j int) bool {
	if s[i].Priority != s[j].Priority {
		return s[i].Priority > s[j].Priority
	} else {
		return s[i].Id < s[i].Id
	}
}

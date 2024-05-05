package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Record struct {
	Id                     int
	Time                   int
	Distance               int
	ActualDeviceNumber     string
	RegisteredDeviceNumber string
	Origin                 string
}

var n int
var records []string
var wrongRecords []string

func main() {
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var record string
		fmt.Scan(&record)
		records = append(records, record)
	}
	records2 := []*Record{}
	for _, record := range records {
		strs := strings.Split(record, ",")
		r := &Record{ActualDeviceNumber: strs[3], RegisteredDeviceNumber: strs[4]}
		r.Id, _ = strconv.Atoi(strs[0])
		r.Time, _ = strconv.Atoi(strs[1])
		r.Distance, _ = strconv.Atoi(strs[2])
		r.Origin = record
		records2 = append(records2, r)
	}
	for _, r1 := range records2 {
		if r1.ActualDeviceNumber != r1.RegisteredDeviceNumber {
			wrongRecords = append(wrongRecords, r1.Origin)
		} else {
			for _, r2 := range records2 {
				if r1 == r2 {
					continue
				}
				if math.Abs(float64(r1.Time-r2.Time)) < 60 && math.Abs(float64(r1.Distance-r2.Distance)) > 5 {
					wrongRecords = append(wrongRecords, r1.Origin)
					break
				}
			}
		}
	}
	if len(wrongRecords) > 0 {
		for i := 0; i < len(wrongRecords)-1; i++ {
			fmt.Printf("%s;", wrongRecords[i])
		}
		fmt.Println(wrongRecords[len(wrongRecords)-1])
	} else {
		fmt.Println("null")
	}
}

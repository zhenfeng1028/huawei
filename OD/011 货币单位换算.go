package main

import (
	"fmt"
	"strconv"
)

var hashmap = map[string]float64{
	"CNY": 100,
	"HKD": float64(100) / float64(123) * 100,
	"JPY": float64(100) / float64(1825) * 100,
	"EUR": float64(100) / float64(14) * 100,
	"GBP": float64(100) / float64(12) * 100,
}

var hashmap2 = map[string]string{
	"fen":       "CNY",
	"cents":     "HKD",
	"sen":       "JPY",
	"eurocents": "EUR",
	"pence":     "GBP",
}

type Record struct {
	Yuan     int    `json:"yuan"`
	YuanName string `json:"yuan_name"`
	Fen      int    `json:"fen"`
	FenName  string `json:"fen_name"`
}

var records []string // 输入的货币记录

func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var record string
		fmt.Scan(&record)
		records = append(records, record)
	}
	recs := []*Record{}
	for _, record := range records {
		var yuanPos, fenPos1, fenPos2 int
		for j := 0; j < len(record); j++ {
			if j > 1 {
				if 'A' <= record[j] && record[j] <= 'Z' && '0' <= record[j-1] && record[j-1] <= '9' {
					yuanPos = j // 记录yuan第一个大写字母下标
				} else if '0' <= record[j] && record[j] <= '9' && 'A' <= record[j-1] && record[j-1] <= 'Z' {
					fenPos1 = j //	记录fen第一个数字下标
				} else if 'a' <= record[j] && record[j] <= 'z' && '0' <= record[j-1] && record[j-1] <= '9' {
					fenPos2 = j // 记录fen第一个小写字母下标
				}
			}
		}
		rec := &Record{}
		if yuanPos > 0 {
			yuan, _ := strconv.Atoi(record[:yuanPos])
			rec.Yuan = yuan
			if fenPos1 == 0 { // 只含yuan
				rec.YuanName = record[yuanPos:]
			} else { // 既含yuan又含分
				rec.YuanName = record[yuanPos:fenPos1]
			}
		}
		if fenPos1 == 0 { // 只含fen
			fen, _ := strconv.Atoi(record[0:fenPos2])
			rec.Fen = fen
			rec.FenName = record[fenPos2:]
		} else { // 既含yuan又含fen
			fen, _ := strconv.Atoi(record[fenPos1:fenPos2])
			rec.Fen = fen
			rec.FenName = record[fenPos2:]
		}
		recs = append(recs, rec)
	}
	sum := 0
	for _, rec := range recs {
		yuanName := rec.YuanName
		yuan := float64(rec.Yuan)
		if rec.Yuan > 0 && rec.Fen > 0 {
			yuan += float64(rec.Fen) / 100
		} else if rec.Yuan == 0 && rec.Fen > 0 {
			yuanName = hashmap2[rec.FenName]
			yuan += float64(rec.Fen) / 100
		}
		sum += int(hashmap[yuanName] * yuan)
	}
	fmt.Println(sum)
}

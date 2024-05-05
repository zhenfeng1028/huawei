package main

import "fmt"

var hashmap = map[int]int{
	1:  31,
	3:  31,
	4:  30,
	5:  31,
	6:  30,
	7:  31,
	8:  31,
	9:  30,
	10: 31,
	11: 30,
	12: 31,
}

func main() {
	var year, month, day int
	fmt.Scan(&year, &month, &day)
	var leapYear bool
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		leapYear = true
	}
	if month == 1 {
		fmt.Println(day)
	} else if month == 2 {
		fmt.Println(31 + day)
	} else {
		days := 0
		for i := 1; i < month; i++ {
			if i == 2 && leapYear {
				days += 29
			} else if i == 2 && !leapYear {
				days += 28
			} else {
				days += hashmap[i]
			}
		}
		days += day
		fmt.Println(days)
	}
}

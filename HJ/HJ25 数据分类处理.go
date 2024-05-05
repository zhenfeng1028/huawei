package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var IArr, RArr []string
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	strs1 := strings.Fields(input.Text())
	IArr = strs1[1:]
	input.Scan()
	strs2 := strings.Fields(input.Text())
	RArr = strs2[1:]
	RArrInt := []int{}
	for _, v := range RArr {
		vv, _ := strconv.Atoi(v)
		RArrInt = append(RArrInt, vv)
	}
	sort.Ints(RArrInt)
	RArr2 := []int{}
	RArr2 = append(RArr2, RArrInt[0])
	for i := 1; i < len(RArrInt); i++ {
		if RArrInt[i] == RArr2[len(RArr2)-1] {
			continue
		}
		RArr2 = append(RArr2, RArrInt[i])
	}
	res := []string{}
	for _, v := range RArr2 {
		s := strconv.Itoa(v)
		exist := false
		count := 0
		tmp := []string{}
		for i, vv := range IArr {
			if strings.Contains(vv, s) {
				exist = true
				count++
				tmp = append(tmp, strconv.Itoa(i))
				tmp = append(tmp, vv)
			}
		}
		if exist {
			res = append(res, s)
			res = append(res, strconv.Itoa(count))
			res = append(res, tmp...)
		}
	}
	fmt.Printf("%d ", len(res))
	for _, v := range res {
		fmt.Printf("%s ", v)
	}
}

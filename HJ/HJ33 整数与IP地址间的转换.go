package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	ip := input.Text()
	input.Scan()
	dec := input.Text()
	fmt.Println(ip2dec(ip))
	fmt.Println(dec2ip(dec))
}

func ip2dec(ip string) int {
	strs := strings.Split(ip, ".")
	binaries := []string{}
	for _, str := range strs {
		num, _ := strconv.Atoi(str)
		binary := strconv.FormatInt(int64(num), 2)
		complete := ""
		if len(binary) < 8 {
			for i := 0; i < 8-len(binary); i++ {
				complete += "0"
			}
		}
		complete += binary
		binaries = append(binaries, complete)
	}
	last := strings.Join(binaries, "")
	res, _ := strconv.ParseInt(last, 2, 64)
	return int(res)
}

func dec2ip(dec string) string {
	num, _ := strconv.Atoi(dec)
	binary := strconv.FormatInt(int64(num), 2)
	complete := ""
	if len(binary) < 32 {
		for i := 0; i < 32-len(binary); i++ {
			complete += "0"
		}
	}
	complete += binary
	binaries := []string{}
	for i := 0; i < 4; i++ {
		binaries = append(binaries, complete[i*8:(i+1)*8])
	}
	nums := []string{}
	for _, binary := range binaries {
		num, _ := strconv.ParseInt(binary, 2, 64)
		nums = append(nums, strconv.Itoa(int(num)))
	}
	res := strings.Join(nums, ".")
	return res
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var ip1, ip2, mask string
	var ipNum1, ipNum2, maskNum int
	var binIp1, binIp2, binMask string
	var max1Idx, min0Idx int
	var strs = []string{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		strs = append(strs, scanner.Text())
	}
	mask, ip1, ip2 = strs[0], strs[1], strs[2]
	ipNum1 = ipStringToByteSlice(ip1)
	ipNum2 = ipStringToByteSlice(ip2)
	maskNum = ipStringToByteSlice(mask)
	if ipNum1 == -1 || ipNum2 == -1 || maskNum == -1 {
		fmt.Println(1)
		return
	}

	max1Idx = -1
	min0Idx = -1
	binMask = strconv.FormatInt(int64(maskNum), 2)
	if len(binMask) < 32 {
		fmt.Println(1)
		return
	}
	for i := 0; i < len(binMask); i++ { // 找1最后位置 0最前位置
		if binMask[i] == '1' {
			max1Idx = i
		}
		if binMask[i] == '0' && min0Idx == -1 {
			min0Idx = i
		}
	}
	if max1Idx == len(binMask)-1 || min0Idx == 0 || max1Idx > min0Idx { // 全是1 或者0出现在开头 或者0后面还有1
		fmt.Println(1)
		return
	}

	binIp1 = strconv.FormatInt(int64(ipNum1), 2)
	binIp2 = strconv.FormatInt(int64(ipNum2), 2)
	res1, res2 := []byte{}, []byte{}
	maskIdx := len(binMask) - 1
	for i := len(binIp1) - 1; i >= 0; i-- {
		res1 = append(res1, binIp1[i]&binMask[maskIdx])
		maskIdx--
	}
	maskIdx = len(binMask) - 1
	for i := len(binIp2) - 1; i >= 0; i-- {
		res2 = append(res2, binIp2[i]&binMask[maskIdx])
		maskIdx--
	}
	if string(res1) == string(res2) {
		fmt.Println(0)
	} else {
		fmt.Println(2)
	}
}

func ipStringToByteSlice(ip string) int {
	s := strings.Split(ip, ".")
	if len(s) != 4 {
		return -1
	}
	r := make([]int, 4)
	var num int
	var err error
	for i := 0; i < 4; i++ {
		num, err = strconv.Atoi(s[i])
		if err != nil {
			return -1
		}
		if num < 0 || num > 255 {
			return -1
		}
		r[i] = num
	}
	return r[0]<<24 + r[1]<<16 + r[2]<<8 + r[3]
}

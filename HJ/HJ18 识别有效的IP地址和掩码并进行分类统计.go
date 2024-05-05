package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 子网掩码为二进制下前面是连续的1，然后全是0

const (
	typeAStart = 1 << 24
	typeAEnd   = 126<<24 + 255<<16 + 255<<8 + 255 // A类地址：1.0.0.0~126.255.255.255
	typeBStart = 128 << 24
	typeBEnd   = 191<<24 + 255<<16 + 255<<8 + 255 // B类地址：128.0.0.0~191.255.255.255
	typeCStart = 192 << 24
	typeCEnd   = 223<<24 + 255<<16 + 255<<8 + 255 // C类地址：192.0.0.0~223.255.255.255
	typeDStart = 224 << 24
	typeDEnd   = 239<<24 + 255<<16 + 255<<8 + 255 // D类地址：224.0.0.0~239.255.255.255
	typeEStart = 240 << 24
	typeEEnd   = 255<<24 + 255<<16 + 255<<8 + 255 // E类地址：240.0.0.0~255.255.255.255

	// 私网ip
	typePrivate1Start = 10 << 24
	typePrivate1End   = 10<<24 + 255<<16 + 255<<8 + 255 // 10.0.0.0~10.255.255.255
	typePrivate2Start = 172<<24 + 16<<16
	typePrivate2End   = 172<<24 + 31<<16 + 255<<8 + 255 // 172.16.0.0~172.31.255.255
	typePrivate3Start = 192<<24 + 168<<16
	typePrivate3End   = 192<<24 + 168<<16 + 255<<8 + 255 // 192.168.0.0~192.168.255.255
)

func main() {
	var ip, mask string
	var lineParts []string
	var a, b, c, d, e, errCount, private int
	var ipNum, maskNum int
	var binMask string
	var max1Idx, min0Idx int

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() { // 使用 ctrl+D 退出循环
		lineParts = strings.Split(scanner.Text(), "~")
		ip = lineParts[0]
		mask = lineParts[1]

		ipNum = ipStringToByteSlice(ip)
		maskNum = ipStringToByteSlice(mask)

		if ipNum>>24 == 0 || ipNum>>24 == 127 {
			continue // 类似于【0.*.*.*】和【127.*.*.*】的IP地址不属于上述输入的任意一类，也不属于不合法ip地址，计数时请忽略
		}

		if ipNum == -1 || maskNum == -1 {
			errCount++ // ip 或 mask 转数字失败
			continue
		}

		max1Idx = -1
		min0Idx = -1
		binMask = strconv.FormatInt(int64(maskNum), 2)
		for i := 0; i < len(binMask); i++ { // 找1最后位置 0最前位置
			if binMask[i] == '1' {
				max1Idx = i
			}
			if binMask[i] == '0' && min0Idx == -1 {
				min0Idx = i
			}
		}

		if max1Idx == len(binMask)-1 || min0Idx == 0 || max1Idx > min0Idx { // 全是1 或者0出现在开头 或者0后面还有1
			errCount++
			continue
		}

		if typePrivate1Start <= ipNum && ipNum <= typePrivate1End ||
			typePrivate2Start <= ipNum && ipNum <= typePrivate2End ||
			typePrivate3Start <= ipNum && ipNum <= typePrivate3End {
			private++
		}

		if typeAStart <= ipNum && ipNum <= typeAEnd {
			a++
		} else if typeBStart <= ipNum && ipNum <= typeBEnd {
			b++
		} else if typeCStart <= ipNum && ipNum <= typeCEnd {
			c++
		} else if typeDStart <= ipNum && ipNum <= typeDEnd {
			d++
		} else if typeEStart <= ipNum && ipNum <= typeEEnd {
			e++
		}
	}
	fmt.Println(a, b, c, d, e, errCount, private)
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
		r[i] = num
	}
	return r[0]<<24 + r[1]<<16 + r[2]<<8 + r[3]
}

package main

import (
	"fmt"
	"math"
)

var pixels []int // 输入像素值

func main() {
	for {
		var pixel int
		n, _ := fmt.Scan(&pixel)
		if n == 0 {
			break
		}
		pixels = append(pixels, pixel)
	}
	target := len(pixels) * 128
	res := 0
	sum := 0
	for k := -127; k <= 128; k++ { // 255-127=128 0+128=128 所以k的范围是-127~128
		var tmp int
		for i := 0; i < len(pixels); i++ {
			newPixel := normalize(pixels[i] + k)
			tmp += newPixel
		}
		if k == -127 || math.Abs(float64(tmp-target)) < math.Abs(float64(sum-target)) {
			sum = tmp
			res = k
		}
	}
	fmt.Println(res)
}

func normalize(pixel int) int {
	if pixel < 0 {
		return 0
	}
	if pixel > 255 {
		return 255
	}
	return pixel
}

package main

import (
	"fmt"
	"sort"
)

type Result struct {
	Price int // 优惠后价格
	Num   int // 使用优惠券数量
}

var moneyOff int     // 满减
var discount int     // 折扣
var noLimit int      // 无门槛
var n int            // 人数
var totalPrice []int // 总价

func main() {
	fmt.Scan(&moneyOff, &discount, &noLimit, &n)
	for i := 0; i < n; i++ {
		var tmp int
		fmt.Scan(&tmp)
		totalPrice = append(totalPrice, tmp)
	}
	for _, tp := range totalPrice {
		p1, num1 := combination1(tp)
		p2, num2 := combination2(tp)
		p3, num3 := combination3(tp)
		p4, num4 := combination4(tp)
		p5, num5 := combination5(tp)
		p6, num6 := combination6(tp)
		res := []*Result{
			{p1, num1},
			{p2, num2},
			{p3, num3},
			{p4, num4},
			{p5, num5},
			{p6, num6},
		}
		sort.Sort(ResultSlice(res))
		fmt.Printf("%d %d\n", res[0].Price, res[0].Num)
	}
}

type ResultSlice []*Result

func (s ResultSlice) Len() int      { return len(s) }
func (s ResultSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s ResultSlice) Less(i, j int) bool {
	if s[i].Price != s[j].Price {
		return s[i].Price < s[j].Price
	} else { // 价格相等时按照优惠券数量排序
		return s[i].Num < s[j].Num
	}

}

// 组合一：先满减再折扣
func combination1(price int) (int, int) {
	price, num1 := MoneyOff(price)
	price, num2 := Discount(price)
	return price, num1 + num2
}

// 组合二：先折扣再满减
func combination2(price int) (int, int) {
	price, num1 := Discount(price)
	price, num2 := MoneyOff(price)
	return price, num1 + num2
}

// 组合三：先满减再无门槛
func combination3(price int) (int, int) {
	price, num1 := MoneyOff(price)
	price, num2 := NoLimit(price)
	return price, num1 + num2
}

// 组合四：先无门槛再满减
func combination4(price int) (int, int) {
	price, num1 := NoLimit(price)
	price, num2 := MoneyOff(price)
	return price, num1 + num2
}

// 组合五：先折扣再无门槛
func combination5(price int) (int, int) {
	price, num1 := Discount(price)
	price, num2 := NoLimit(price)
	return price, num1 + num2
}

// 组合六：先无门槛再折扣
func combination6(price int) (int, int) {
	price, num1 := NoLimit(price)
	price, num2 := Discount(price)
	return price, num1 + num2
}

// 满减
func MoneyOff(price int) (int, int) {
	var num int
	if price/100 >= 1 && moneyOff > 0 {
		if price/100 >= moneyOff {
			num += moneyOff // 满减全部使用
			price -= 10 * moneyOff
		} else {
			num += price / 100 // 使用部分满减
			price -= 10 * (price / 100)
		}
	}
	return price, num
}

// 折扣
func Discount(price int) (int, int) {
	var num int
	if discount > 0 {
		num++
		price = int(0.92 * float64(price))
	}
	return price, num
}

// 无门槛
func NoLimit(price int) (int, int) {
	count := noLimit
	if noLimit > 0 {
		for price >= 0 && count > 0 {
			price -= 5
			count--
		}
	}
	return price, noLimit - count
}

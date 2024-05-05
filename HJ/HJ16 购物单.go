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
	str := strings.Split(input.Text(), " ")
	money, _ := strconv.Atoi(str[0])
	n, _ := strconv.Atoi(str[1])
	items := make([]item, n)
	for i := 0; i < n; i++ {
		input.Scan()
		it := strings.Split(input.Text(), " ")
		v, _ := strconv.Atoi(it[0])
		p, _ := strconv.Atoi(it[1])
		q, _ := strconv.Atoi(it[2])
		items[i] = item{
			v:    v,
			p:    p,
			q:    q,
			acc1: nil,
			acc2: nil,
		}
	}
	// 将附件依附到主件
	for i, k := range items {
		if k.q != 0 {
			if items[k.q-1].acc1 == nil {
				items[k.q-1].acc1 = &items[i]
			} else {
				items[k.q-1].acc2 = &items[i]
			}
		}
	}
	// 由于每个东西只能买一件，并且买了主件才能买附件，同时主件的附件数量是确定的不大于2
	// 因此我们可以看成购买就是针对于主件，只不过主件的附件数量不同而已，因为只能买一次
	// 因此主件和附件的搭配只能选一种，之后就可以看成01背包问题，只不过背包的物品是多选一的
	arr := make([][]int, n+1)
	for i, _ := range arr {
		arr[i] = make([]int, money+1)
	}
	cnt := 1
	for i := 0; i < n; i++ {
		// 附件直接跳过
		if items[i].q != 0 {
			continue
		}
		// 构造各个选择的价格和满意度
		var (
			v0   = items[i].v // 只有主件
			myd0 = v0 * items[i].p
			v1   = v0 // 主件加附件1
			myd1 = myd0
			v2   = v0 // 主件加附件2
			myd2 = myd0
			v3   = v0 // 主件加两个附件
			myd3 = myd0
		)
		if items[i].acc1 != nil {
			v1 += items[i].acc1.v
			myd1 += items[i].acc1.v * items[i].acc1.p
		}
		if items[i].acc2 != nil {
			v2 += items[i].acc2.v
			myd2 += items[i].acc2.v * items[i].acc2.p
		}
		if items[i].acc1 != nil && items[i].acc2 != nil {
			v3 = v3 + items[i].acc1.v + items[i].acc2.v
			myd3 = myd3 + items[i].acc1.v*items[i].acc1.p + items[i].acc2.v*items[i].acc2.p
		}
		for j := 1; j <= money; j++ {
			arr[cnt][j] = arr[cnt-1][j]
			if j >= v0 {
				arr[cnt][j] = max(arr[cnt][j], arr[cnt-1][j-v0]+myd0)
			}
			if j >= v1 && v1 > v0 {
				arr[cnt][j] = max(arr[cnt][j], arr[cnt-1][j-v1]+myd1)
			}
			if j >= v2 && v2 > v0 {
				arr[cnt][j] = max(arr[cnt][j], arr[cnt-1][j-v2]+myd2)
			}
			if j >= v3 && v3 > v0 {
				arr[cnt][j] = max(arr[cnt][j], arr[cnt-1][j-v3]+myd3)
			}
		}
		cnt++
	}
	fmt.Println(arr[cnt-1][money])
}

type item struct {
	v    int
	p    int
	q    int
	acc1 *item
	acc2 *item
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

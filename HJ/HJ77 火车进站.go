package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var N int
	fmt.Scan(&N)
	trains := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&trains[i])
	}

	var in []int      // 入站序列
	var out []int     // 出站序列
	var res [][]int   // 保存结果
	var dfs func(int) // 递归函数

	dfs = func(index int) {
		if index == N && len(in) == 0 { // 递归边界
			tmp := []int{}
			for _, v := range out {
				tmp = append(tmp, v)
			}
			res = append(res, tmp)
			// res = append(res, out)	// 这种方式存在问题，res引用了out的底层数组，当out变化时res也发生变化
			return
		}
		// 进
		if index < N {
			in = append(in, trains[index])
			dfs(index + 1)
			in = in[:len(in)-1] // 回溯
		}
		// 出
		if len(in) != 0 {
			// 栈顶元素
			top := in[len(in)-1]
			out = append(out, top)
			in = in[:len(in)-1]
			dfs(index)
			in = append(in, top)   // 回溯
			out = out[:len(out)-1] // 回溯
		}
	}

	dfs(0)

	strs := []string{}
	for _, v := range res {
		str := ""
		for _, vv := range v {
			str += strconv.Itoa(vv)
			str += " "
		}
		strs = append(strs, str)
	}
	sort.Strings(strs)
	for _, str := range strs {
		fmt.Println(str)
	}
}

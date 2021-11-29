package main

import (
	"fmt"
	"sort"
)

type twonum struct {
	x int
	y int
}

func main() {
	// 给你一个按递增顺序排序的数组 arr 和一个整数 k 。数组 arr 由 1 和若干 素数组成，且其中所有整数互不相同。

	// 对于每对满足 0 < i < j < arr.length 的 i 和 j ，可以得到分数 arr[i] / arr[j] 。

	// 那么第k个最小的分数是多少呢? 以长度为 2 的整数数组返回你的答案, 这里answer[0] == arr[i]且answer[1] == arr[j] 。

	// 来源: 力扣（LeetCode）
	// 链接: https://leetcode-cn.com/problems/k-th-smallest-prime-fraction
	// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

	// 题解 暴力 选最小的
	arr := []int{1, 2, 3, 5}
	k := 3
	res := make([]twonum, 0)
	for i, v1 := range arr {
		for _, v2 := range arr[i+1:] {
			res = append(res, twonum{x: v1, y: v2})
		}
	}

	sort.Slice(res, func(i, j int) bool {
		a, b := res[i], res[j]
		return a.x*b.y < b.x*a.y
	})
	fmt.Println([]int{res[k-1].x, res[k-1].y})
}

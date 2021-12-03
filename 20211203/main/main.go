package main

import (
	"fmt"
	"math/rand"
	"time"
)

func make_arr(length int) []int {
	arr := []int{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		arr = append(arr, rand.Intn(101))
	}
	return arr
}

func main() {
	// 	给你一个整数数组 nums 和一个整数 k ，按以下方法修改该数组：

	// 选择某个下标 i 并将 nums[i] 替换为 -nums[i] 。
	// 重复这个过程恰好 k 次。可以多次选择同一个下标 i 。

	// 以这种方式修改数组后，返回数组 可能的最大和 。

	// 来源：力扣（LeetCode）
	// 链接：https://leetcode-cn.com/problems/maximize-sum-of-array-after-k-negations
	// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
	arr := make_arr(10)
	fmt.Println(arr)
	k := 5
	nummap := make(map[int]int)
	sum := 0

	for _, num := range arr {
		nummap[num]++
		sum += num
	}

	for i := -100; k != 0 && i <= 0; i++ {
		if nummap[i] != 0 {
			if k-nummap[i] >= 0 {
				sum += nummap[i] * (-i) * 2
				k -= nummap[i]
			} else {
				sum += k * (-i) * 2
				k = 0
			}
		}
	}

	if k != 0 && k%2 == 1 && nummap[0] == 0 {
		for i := 1; i <= 100; i++ {
			if nummap[i] > 0 {
				sum -= i * 2
				break
			}
		}
	}

	fmt.Println(sum)
}

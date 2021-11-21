package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 和谐数组是指一个数组里元素的最大值和最小值之间的差别 正好是 1 。

// 现在，给你一个整数数组 nums ，请你在所有可能的子序列中找到最长的和谐子序列的长度。

// 数组的子序列是一个由数组派生出来的序列，它可以通过删除一些元素或不删除元素、且不改变其余元素的顺序而得到。

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/longest-harmonious-subsequence
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func make_arr(length int) []int {
	arr := []int{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		arr = append(arr, rand.Intn(10))
	}
	return arr
}

func main() {
	fmt.Println("数组长度为：")
	var length int
	fmt.Scanln(&length)

	arr := make_arr(length)

	fmt.Println("数组为: ", arr)

	// 处理

	ml := make(map[int]int)

	for _, v := range arr {
		ml[v] += 1
	}

	sl := 0

	for i, v := range ml {
		if c1 := ml[i+1]; c1 > 0 && v+c1 > sl {
			sl = v + c1
		}
	}

	fmt.Println("题解为:", sl)

}

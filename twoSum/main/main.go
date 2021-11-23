package main

import "fmt"

func twoSum() []int {
	arr := [...]int{3, 2, 4}
	target := 6

	if len(arr) == 2 {
		return []int{0, 1}
	}

	intmap := make(map[int]int)
	for i, v := range arr {
		intmap[v] = i // 存储最后一个下标
	}

	for i, v := range arr {
		v, ok := intmap[target-v]
		if ok {
			if i != v {
				return []int{i, v}
			}
		}
	}
	return []int{}

	// 标准答案

	// hashTable := map[int]int{}
	// for i, x := range nums {
	//     if p, ok := hashTable[target-x]; ok {
	//         return []int{p, i}
	//     }
	//     hashTable[x] = i
	// }
	// return nil

}

func main() {
	// 两数之和
	arr := twoSum()
	fmt.Println(arr)
}

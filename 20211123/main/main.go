package main

import (
	"fmt"
)

func jg() bool {
	s := "aa"
	goal := "aa"
	strmap1 := make(map[int]int, 26)
	strmap2 := make(map[int]int, 26)
	flag := 0

	if len(s) != len(goal) {
		return false
	}

	for i := range s {
		a := int(s[i] - byte('a'))
		b := int(goal[i] - byte('a'))
		strmap1[a]++
		strmap2[b]++
		if a != b {
			flag++
		}
	}

	bool_ := false

	for i := 0; i < 26; i++ {
		if strmap1[i] != strmap2[i] {
			return false
		}
		if strmap1[i] > 1 {
			bool_ = true
		}
	}

	return flag == 2 || (flag == 0 && bool_)

}

func main() {
	if jg() {
		fmt.Println("对")
	} else {
		fmt.Println("错")
	}
}

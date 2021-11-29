package main

import (
	"fmt"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return nil
}

func main() {
	// var ln1 *ListNode
	// var ln2 *ListNode
	// res := addTwoNumbers(ln1, ln2)
	// fmt.Printf("%v\n", res)
	num := 807
	n := 10
	for {
		temp := num / n * n
		if temp != 0 {
			fmt.Println(num - temp)
			num /= 10
		} else {
			if num/n != 0 {
				fmt.Println(num - temp)
				num /= 10
			} else {
				fmt.Println(num - temp)
				break
			}
		}
	}

}

package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	for {
		if root.Val == val {
			return root
		}
		if root.Val > val {
			if root.Left != nil {
				root = root.Left
			} else {
				return nil
			}
		}
		if root.Val < val {
			if root.Right != nil {
				root = root.Right
			} else {
				return nil
			}
		}
	}
}

func main() {
	var tn *TreeNode
	res := searchBST(tn, 63)
	fmt.Printf("%v\n", res)
}

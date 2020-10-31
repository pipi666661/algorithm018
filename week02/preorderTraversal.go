package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func preorderTraversal(root *TreeNode) (res []int) {
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		res = append(res, node.Val)
		preorder((node.Left))
		preorder(node.Right)
	}
	preorder(root)
	return
}


func PrintArr(arr []int){

	for _,v := range arr{
		fmt.Printf("%d ",v)
	}
	fmt.Println()
}


func main(){
	TreeNode3 := TreeNode{
		3,
		nil,
		nil,
	}

	TreeNode2 := TreeNode{
		2,
		&TreeNode3,
		nil,
	}

	TreeNode := TreeNode{
		1,
		nil,
		&TreeNode2,
	}
	PrintArr(preorderTraversal(&TreeNode))
}
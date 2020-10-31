package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	 Right *TreeNode
}


func inorderTraversal(root *TreeNode) (res []int) {
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
	}
	inorder(root)
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
	PrintArr(inorderTraversal(&TreeNode))
}
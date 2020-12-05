package main

import "fmt"

type Node struct{
	Val int
	Children []*Node
}


func PrintArr2(arr [][]int){

	for _,v := range arr{

		for _,k := range v{
			fmt.Printf("%d ",k)
		}

	}
	fmt.Println()
}


func levelOrder(root *Node) [][]int {
	result := make([][]int, 0, 1000)
	if root == nil {
		return result
	}
	q := []*Node{root}
	for len(q) > 0 {
		size := len(q)
		item := make([]int, 0, size)
		for i := 0; i < size; i++ {
			item = append(item, q[i].Val)
			for _, child := range q[i].Children {
				q = append(q, child)
			}
		}
		result = append(result, item)
		q = q[size:]
	}
	return result
}


func main(){
	Node5 := Node{
		5,
		nil,

	}

	Node6 := Node{
		6,
		nil,

	}

	Node3 :=Node{
		3,
		[]*Node{&Node5,&Node6},
	}

	Node2 :=Node{
		2,
		nil,
	}
	Node4 :=Node{
		4,
		nil,
	}
	Node1 :=Node{
		1,
		[]*Node{&Node3,&Node2,&Node4},
	}

	PrintArr2(levelOrder(&Node1))
}
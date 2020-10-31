package main

import "fmt"

type Node struct{
	Val int
	Children []*Node
}


func preorder(root *Node) []int {

	var result []int
	result = transorder(root,result)
	return result


}




func transorder(root *Node, data []int) []int{
	if nil != root{
		data = append(data, root.Val)
			fmt.Println(root.Val)
	}else{
		return data

	}

	if nil != root.Children{
	//	fmt.Println("len = ",len(root.Children) )
		for i:=0; i<len(root.Children); i++{
			data = transorder(root.Children[i],data)
		}
	}

	return data




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

	preorder(&Node1)

}
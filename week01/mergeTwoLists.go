package main

import "fmt"

type ListNode struct {
	     Val int
	     Next *ListNode
	}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 ==nil {
		return l2
	}

	if l2 ==nil {
		return l1
	}

	returnlist := &ListNode{}
	prehead :=returnlist

	for l1 !=nil && l2!=nil{
		if l1.Val < l2.Val{
			prehead.Next = l1
			l1 = l1.Next
		}else{
			prehead.Next = l2
			l2 = l2.Next
		}
		prehead = prehead.Next
	}

	if l1!=nil{
		prehead.Next=l1
	}

	if l2!=nil{
		prehead.Next=l2
	}

	return returnlist.Next


}

func main(){

	ListNode4 := &ListNode{4,nil}
	ListNode3 := &ListNode{3,ListNode4}

	ListNode2 := &ListNode{2,ListNode3}
	ListNode1 := &ListNode{1,ListNode2}
	ListNode8 :=  &ListNode{3,nil}

	ListNode7 :=  &ListNode{2,ListNode8}


	ListNode6 :=  &ListNode{1,ListNode7}
	ListNode5 :=  &ListNode{1,ListNode6}

	printList(ListNode1)
	printList(ListNode5)

	printList(mergeTwoLists(ListNode1,ListNode5))

}

func printList(l1 *ListNode){
	p :=l1
	for p.Next != nil{
		fmt.Printf("%d ",p.Val)
		p = p.Next
	}
	fmt.Println()

}
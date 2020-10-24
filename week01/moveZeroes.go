package main

import "fmt"

func PrintArr(arr []int){

	for _,v := range arr{
		fmt.Printf("%d ",v)
	}
	fmt.Println()
}


func main(){
	arr := []int{0,1,0,3,12}
	PrintArr(arr)
	moveZeroes(arr)
	PrintArr(arr)
}


func moveZeroes(nums []int)  {


	i := 0

	for _,n := range nums{
		if n != 0{
			nums[i] = n
			i ++
		}
	}

	for ;i < len(nums); i++{
		nums[i] = 0
	}




}
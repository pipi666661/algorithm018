package main

import "fmt"

func PrintArr(arr []int){

	for _,v := range arr{
		fmt.Printf("%d ",v)
	}
	fmt.Println()
}



func main(){

	arr := []int{1,1,2,3,3}
	PrintArr(arr)

	fmt.Println(removeDuplicates(arr))
}

func removeDuplicates(nums []int) int {
	j := 0
	for i := 1; i < len(nums); i++ {
		if nums[j] != nums[i] {
			j++
			nums[j] = nums[i]
		}
	}
	return j + 1

}
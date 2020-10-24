package main

import "fmt"


func twoSum(nums []int, target int) []int {

	length := len(nums)

	var result []int

	for i:= 0 ; i< length;i++{
		x := target - nums[i]

		for j:= i+1 ; j< length; j++{
			if x == nums[j]{
				result = append(result,i)
				result = append(result,j)
			}

		}

	}
	return result

}

func main(){
	nums :=[]int{2,7,9,11}
	PrintArr(twoSum(nums,9))

}
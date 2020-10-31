package main

import "fmt"

func PrintArr(arr []int){

	for _,v := range arr{
		fmt.Printf("%d ",v)
	}
	fmt.Println()
}


func twoSum(nums []int, target int) []int {

	hashmap := make(map[int]int)


	for i,x := range nums{

		if p,ok := hashmap[target - x]; ok{

			return []int{p,i}

		}else{
			hashmap[x] = i
		}

	}
	return nil

}

func main(){
	nums :=[]int{2,7,9,11}
	PrintArr(twoSum(nums,9))

}
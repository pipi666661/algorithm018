package main

import "fmt"

func PrintArr(arr []int){

	for _,v := range arr{
		fmt.Printf("%d ",v)
	}
	fmt.Println()
}



func main(){
	digits :=[]int{1,2,4}
	arr :=plusOne(digits)


	PrintArr(arr)


}

func plusOne(digits []int) []int {
	length :=len(digits)
	for i := length-1; i>=0; i--{
		if digits[i] != 9{
			digits[i]++
			return digits
		}else{
			digits[i] = 0
		}
	}

	var tmp = make([]int,1)
	tmp[0]=1
	tmp = append(tmp,digits...)
	return tmp

}
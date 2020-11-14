package main

import "fmt"

func main(){

	fmt.Println(lemonadeChange([]int{5,5,5,10,20}))

}


func lemonadeChange(bills []int) bool {


	has := make(map[int]int)

	for _,bill := range bills{

		if has[5] < 0 || has[10] <0{
			return false
		}

		if bill == 5{
			has[5]++
		}
		if bill == 10{

			has[5]--
			has[10]++
		}

		if bill == 20{

			if has[5] > 0 && has[10] > 0{
				has[5]--
				has[10]--
				has[20]++
			}else if has[5] > 2{
				has[5]=has[5] -3
				has[20]++

			}else{
				return false
			}

		}



	}
	return true
}

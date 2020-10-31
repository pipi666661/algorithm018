package main

import (
	"fmt"
	"sort"
	"strings"
)

func PrintArr2(arr [][]string){

	for _,v := range arr{

		for _,k := range v{
			fmt.Printf("%s ",k)
		}

	}
	fmt.Println()
}

func groupAnagrams(strs []string) [][]string {
	group := make(map[string]int)
	result := make([][]string,0); //保存结果
	index := 0;
	for _ , v := range strs {
		vArr := strings.Split(v,"")
		sort.Strings(vArr)
		toStr := strings.Join(vArr,"");
		if idx, ok := group[toStr];!ok{
			group[toStr] = index
			result = append(result,[]string{v})
			index++
		}else {
			result[idx] = append(result[idx], v)
		}
	}
	return result
}


func main(){

	PrintArr2(groupAnagrams([]string{"ret","ter","abc","etr"}))




}
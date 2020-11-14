package main

import (
	"fmt"
	"sort"
)


func main(){
	g := []int{1,2,3}
	s := []int{1,1}
	fmt.Printf("%v",findContentChildren(g,s))
}

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	i, j := 0, 0
	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			i++
		}
		j++
	}
	return i


}
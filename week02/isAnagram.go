package main

import "fmt"


func isAnagram(s string, t string) bool {

	if len(s) != len(t){
		return false
	}
	//创建两个仿bitmap的数组
	wordstable_s :=  [26]int{}
	wordstable_t :=  [26]int{}
	for i:=0;i<len(s);i++{
		//使用asc2码 计算对应仿bitmap数组的下标
		index := s[i] - 'a'
		wordstable_s[index] ++
	}
	for i:=0;i<len(t);i++{
		index := t[i] - 'a'
		wordstable_t[index] ++
	}
	//最后比较一下两个数组就好了
	return wordstable_s == wordstable_t

}

func main(){
	s := "anagram"
	t := "nagaram"

	fmt.Println(isAnagram(s, t))

}
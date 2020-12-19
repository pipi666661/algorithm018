package week09
func reverseWords(s string) string {
	arr := []byte(s)
	start := 0
	end := 0
	flag := false
	for i := 0; i < len(arr); i ++{
		if string(arr[i]) == " " {
			flag = true
			end = i
			rev(arr, start, end-1)
			start = i+1
		}
	}
	if flag {
		rev(arr, end+1, len(arr)-1)
	}else {
		rev(arr, 0, len(arr)-1)
	}
	return string(arr)
}

func rev(arr []byte,left, right int)  {
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left ++
		right --
	}
}


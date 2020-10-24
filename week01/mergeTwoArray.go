package main

func merge(nums1 []int, m int, nums2 []int, n int) {

	for m > 0 && n > 0 {
		if nums1[m-1] > nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}
	if m == 0 {
		for i := 0; i < n; i++ {
			nums1[i] = nums2[i]
		}
	}

	//return nums1
}


func main(){
	nums1 := []int{1,4,6,0,0,0}
	nums2 := []int{3,5,8}


	PrintArr(nums1)
	PrintArr(nums2)

	merge(nums1,3,nums2,3)
	PrintArr(nums1)
}



package main

import "fmt"

func partition(nums []int, lo, hi int) int {
	pivot := nums[hi]
	i := lo - 1
	for j := lo; j < hi; j++ {
		if nums[j] > pivot {
			i ++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i + 1], nums[hi] = nums[hi], nums[i+1]
	return i + 1
}
func findKthLargest(nums []int, k int) int {
	lastE := nums[len(nums) -1]
	lastPos := partition(nums, 0, len(nums)-1)
	if lastPos == k - 1 {
		return lastE
	}else if lastPos > k - 1 {
		return findKthLargest(nums[:lastPos], k)
	} else {
		return findKthLargest(nums[lastPos + 1:], k-1-lastPos)
	}
}

func main() {
	nums := []int{0, 2, -1}
	fmt.Println(findKthLargest(nums, 2))
}
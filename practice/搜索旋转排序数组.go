// 解题文章：https://blog.csdn.net/tzh_linux/article/details/103997885
package main

import (
	"fmt"
)

func search (nums []int, target int) int {
	nlen := len(nums)
	if nlen == 0 {
		return -1
	}
	left := 0
	right := nlen - 1
	for left <= right {
		mid := (right - left) / 2 + left
		if nums[mid] == target {
			return mid
		}
		// 先找出在左区间的各种情况
		if nums[left] < nums[mid] && nums[left] <= target && target < nums[mid] {
			right = mid - 1
		} else if nums[left] > nums[mid] && target < nums[mid] {
			right = mid - 1
		} else if nums[left] > nums[mid] && target >= nums[left] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func main(){
	nums := []int{3, 5, 1}
	idx := search(nums, 3)
	fmt.Println(idx)
}

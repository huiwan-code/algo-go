package algo


func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	p := nums[right]
	pos := left
	for i := left; i < right; i ++ {
		if nums[i] < p {
			nums[i], nums[pos] = nums[pos], nums[i]
			pos++
		}
	}
	nums[pos], nums[right] = nums[right], nums[pos]
	quickSort(nums, left, pos - 1)
	quickSort(nums, pos + 1, right)
}

// 归并排序
func mergeSort(nums []int) []int{
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2
	left := mergeSort(nums[0:mid])
	right := mergeSort(nums[mid:])
	result := merge(left, right)
	return result
}

func merge(left, right []int) []int{
	l := 0
	r := 0
	result := make([]int, 0)
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return result
}

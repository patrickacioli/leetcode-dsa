package main

func sortedSquares(nums []int) []int {
	size := len(nums)
	response := make([]int, size)
	left := 0
	right := size - 1
	for i := 0; i < size; i++ {
		if nums[right]*nums[right] >= nums[left]*nums[left] {
			response[size-1-i] = nums[right] * nums[right]
			right--
		} else {
			response[size-1-i] = nums[left] * nums[left]
			left++
		}
	}
	return response
}

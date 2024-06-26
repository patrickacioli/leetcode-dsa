package main

// You are given an integer array nums consisting of n elements, and an integer k.
// Find a contiguous subarray whose length is equal to k that has the maximum average value and return this value.
// Any answer with a calculation error less than 10-5 will be accepted.
func findMaxAverage(nums []int, k int) float64 {
	ans := 0
	for _, i := range nums[0:k] {
		ans += i
	}
	mean := ans

	for i := k; i < len(nums); i++ {
		start := i - k
		end := i
		mean = mean + nums[end] - nums[start]
		if mean > ans {
			ans = mean
		}
	}

	return float64(ans) / float64(k)
}

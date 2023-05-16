package main

import "fmt"

func main() {
	fmt.Printf("%v", summaryRanges([]int{0, 1, 2, 4, 5, 7}))
}

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	left := nums[0]
	res := make([]string, 0, len(nums))
	for i := 1; i < len(nums); i++ {
		println(nums[i-1] + 1)
		println(nums[i])
		if nums[i-1]+1 != nums[i] {
			if left != nums[i-1] {
				res = append(res, fmt.Sprintf("%d->%d", left, nums[i-1]))
			} else {
				res = append(res, fmt.Sprintf("%d", left))
			}
			left = nums[i]
		}
	}
	return res
}

package main

func main() {
	println(singleNumber([]int{2, 2, 1}))
}

func singleNumber(nums []int) int {
	for i := 1; i < len(nums); i++ {
		nums[0] ^= nums[i]
	}
	return nums[0]
}

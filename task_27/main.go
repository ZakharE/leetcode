package main

func main() {
	println(removeElement([]int{2, 2}, 2))
}

func removeElement(nums []int, val int) int {
	i, j := 0, len(nums)-1
	res := len(nums)
	for i <= j {
		if nums[j] == val {
			j--
			res--
		} else if nums[i] == val {
			nums[j], nums[i] = nums[i], nums[j]
			i++
			j--
			res--
		} else {
			i++
		}

	}
	return res
}

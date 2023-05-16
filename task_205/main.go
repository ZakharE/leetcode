package main

func main() {
	println(containsNearbyDuplicate([]int{1, 0, 1, 1}, 1))
}

func containsNearbyDuplicate(nums []int, k int) bool {
	idxs := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if v, ok := idxs[nums[i]]; ok {
			if abs(v-i) <= k {
				return true
			}
		}
		idxs[nums[i]] = i
	}

	return false
}
func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

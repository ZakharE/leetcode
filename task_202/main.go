package main

func main() {
	println(isHappy(1))
}
func isHappy(n int) bool {
	m := make(map[int]struct{})
	m[n] = struct{}{}
	for {
		if n == 1 {
			return true
		}
		n = sum(n)
		if _, ok := m[n]; ok {
			return false
		}
		m[n] = struct{}{}

	}
}

func sum(a int) int {
	res := 0
	for a > 0 {
		rem := a % 10
		res += rem * rem
		a = a / 10
	}
	return res
}

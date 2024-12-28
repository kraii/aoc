package maths

func LcmAll(nums []int) int {
	result := 1
	for _, n := range nums {
		result = Lcm(result, n)
	}
	return result
}

func Lcm(a int, b int) int {
	return (a * b) / Gcd(a, b)
}

func Gcd(a int, b int) int {
	if b == 0 {
		return a
	} else {
		return Gcd(b, a%b)
	}
}

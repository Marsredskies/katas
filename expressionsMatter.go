package kata

import "math"

// ExpressionMatter https://www.codewars.com/kata/5ae62fcf252e66d44d00008e/train/go
func ExpressionMatter(a int, b int, c int) int {
	return int(math.Max(math.Max(float64(a*(b+c)), float64(a*b*c)), math.Max(float64(a+b+c), float64((a+b)*c))))
}

package kata

// Summation https://www.codewars.com/kata/55d24f55d7dd296eb9000030/train/go
func Summation(n int) int {
	sum := (n * (n + 1)) / 2 // Gauss's formula
	return sum
}

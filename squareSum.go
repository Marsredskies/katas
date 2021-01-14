// https://www.codewars.com/kata/515e271a311df0350d00000f/train/go
package kata

func SquareSum(numbers []int) int {
	var sum int
	for _, v := range numbers {
		sum += v * v
	}
	return sum
}

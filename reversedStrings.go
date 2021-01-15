package kata

// Solution https://www.codewars.com/kata/5168bb5dfe9a00b126000018/train/go
func Solution(word string) string {
	runes := []rune(word)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)

}

package kata

// Hero https://www.codewars.com/kata/59ca8246d751df55cc00014c/train/go
func Hero(bullets, dragons int) bool {
	if bullets >= dragons*2 {
		return true
	}
	return false

}

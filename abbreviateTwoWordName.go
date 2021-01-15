package kata

import (
	"fmt"
	"strings"
)

// AbbrevName https://www.codewars.com/kata/57eadb7ecd143f4c9c0000a3/train/go
func AbbrevName(name string) string {
	splitted := strings.Split(name, " ")
	first := strings.ToTitle(string([]rune(splitted[0])[0]))
	second := strings.ToTitle(string([]rune(splitted[1])[0]))
	return fmt.Sprintf("%s.%s", first, second)
}

// https://www.codewars.com/kata/5bb904724c47249b10000131/train/go

package kata

import "strings"

func Points(games []string) int {
	allPoints := 0
	for _, match := range games {
		score := strings.Split(match, ":")
		x, y := score[0], score[1]
		switch {
		case x > y:
			allPoints += 3
		case x == y:
			allPoints += 1
		}

	}
	return allPoints

}

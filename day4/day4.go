package main

import "strconv"

func main() {
	min, max := 156218, 652527

	count := 0
outer:
	for i := min; i <= max; i++ {
		s := strconv.Itoa(i)
		double := false
		for i, c := range s[1:] {
			if c < rune(s[i]) {
				continue outer
			}
		}
		for i, c := range s[1:] {
			if c == rune(s[i]) && (i == 0 || c != rune(s[i-1])) && (i == len(s)-2 || c != rune(s[i+2])) {
				// there's one before us
				// make sure there's not two before us
				// make sure there's not one after us
				double = true
			}
		}
		if double {
			count++
		}
	}
	println(count)
}

// time: 00:11:51
// time: 00:18:18

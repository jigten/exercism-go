package isogram

import "unicode"

func IsIsogram(word string) bool {
	seen := map[rune]bool{}

	for _, c := range word {
		if c == ' ' || c == '-' {
			continue
		}

		lower := unicode.ToLower(c)

		if seen[lower] {
			return false
		}
		seen[lower] = true
	}

	return true
}

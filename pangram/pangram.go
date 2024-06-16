package pangram

import (
	"unicode"
)

func IsPangram(input string) bool {
	seen := map[rune]bool{}

	for _, c := range input {
		if unicode.IsLetter(c) {
			s := unicode.ToLower(c)
			_, ok := seen[s]

			if !ok {
				seen[s] = true
			}
		}
	}

	return len(seen) == 26
}

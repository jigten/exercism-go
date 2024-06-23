package isbn

import (
	"strconv"
	"strings"
	"unicode"
)

func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")

	if len(isbn) != 10 {
		return false
	}

	val := 0

	for i, c := range isbn {
		if unicode.IsLetter(c) && c != 'X' {
			return false
		}

		if c == 'X' {
			if i != len(isbn)-1 {
				return false
			}
			val += 10
		}

		ic, _ := strconv.Atoi(string(c))
		val += (10 - i) * ic
	}

	return val%11 == 0
}

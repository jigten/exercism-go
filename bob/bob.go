// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

func isUpperCharacters(s string) bool {
	isAllSpecial := true
	for _, r := range s {
		if unicode.IsLetter(r) && !unicode.IsUpper(r) {
			return false
		}

		if unicode.IsLetter(r) {
			isAllSpecial = false
		}
	}
	return true && !isAllSpecial
}

// Hey should have a comment documenting it.
func Hey(remark string) string {
	trimmedRemark := strings.TrimSpace(remark)

	if trimmedRemark == "" {
		return "Fine. Be that way!"
	}

	isQuestion := trimmedRemark[len(trimmedRemark)-1] == '?'
	isShouted := isUpperCharacters(trimmedRemark)

	if isQuestion && !isShouted {
		return "Sure."
	}

	if !isQuestion && isShouted {
		return "Whoa, chill out!"
	}

	if isQuestion && isShouted {
		return "Calm down, I know what I'm doing!"
	}

	return "Whatever."
}

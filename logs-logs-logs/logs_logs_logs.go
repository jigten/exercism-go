package logs

import (
	"fmt"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	application := "default"

	for _, r := range log {
		code := fmt.Sprintf("%U", r)

		if code == "U+2757" {
			application = "recommendation"
			break
		}

		if code == "U+1F50D" {
			application = "search"
			break
		}

		if code == "U+2600" {
			application = "weather"
			break
		}
	}

	return application
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	replaced := ""

	for _, r := range log {
		if r == oldRune {
			replaced += string(newRune)
		} else {
			replaced += string(r)
		}
	}
	return replaced
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}

package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*=-]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`"(?i).*password.*"`)
	count := 0
	for _, line := range lines {
		if re.MatchString(line) {
			count += 1
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d*`)
	return re.ReplaceAllLiteralString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\w+)`)
	newLines := []string{}

	for _, line := range lines {
		founds := re.FindStringSubmatch(line)
		if founds != nil {
			line = fmt.Sprintf("[USR] %s %s", founds[1], line)
		}
		newLines = append(newLines, line)
	}

	return newLines
}

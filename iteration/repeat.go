package iteration

import "strings"

func Repeat(letter string, count int) string {
	var repeated strings.Builder
	for range count {
		repeated.WriteString(letter)
	}
	return repeated.String()
}

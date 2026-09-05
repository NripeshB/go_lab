package iteration

import "strings"

func Repeat(character string, rep int) string {
	var ans strings.Builder
	for rep > 0 {
		ans.WriteString(character)
		rep--
	}
	return ans.String()
}

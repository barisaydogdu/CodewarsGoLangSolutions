package mumbling

import "strings"

func Accum(s string) string {
	var result strings.Builder
	counter := 0

	for i := 0; i < len(s); i++ {
		//A
		result.WriteString(strings.ToUpper(string(s[i]))) //First letter is Uppercase

		//Other letters repeat with counter and lower case
		result.WriteString(strings.Repeat(strings.ToLower(string(s[i])), counter))

		//add "-" character
		if i != len(s)-1 {
			result.WriteString("-")
		}
		counter++
	}
	return result.String()
}

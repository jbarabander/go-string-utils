package stringUtils

import "unicode"

// Capitalize capitalizes the first letter of a string
func Capitalize(str string) string {
	r := []rune(str)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

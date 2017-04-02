package stringUtils

// Join takes a slice of strings and combines them
func Join(strSlice []string, joiner string) string {
	combinedStr := ""
	lastIndex := len(strSlice) - 1
	for i := 0; i < lastIndex; i++ {
		combinedStr += strSlice[i] + joiner
	}
	combinedStr += strSlice[lastIndex]
	return combinedStr
}

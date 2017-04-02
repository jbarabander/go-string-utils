package stringUtils

// Split splits into a slice of strings based on the string given
func Split(str string, strToSplitOn string) []string {
	r := []rune(str)
	splitSlice := []rune(strToSplitOn)

	splitSliceLen := len(splitSlice)
	currSplitSliceIndex := 0

	var strSlice []string
	var currentStr []rune
	for i := 0; i < len(r); i++ {
		if r[i] == splitSlice[currSplitSliceIndex] {
			currSplitSliceIndex++
		} else {
			currSplitSliceIndex = 0
		}
		if currSplitSliceIndex == splitSliceLen {
			currSplitSliceIndex = 0
			strWithSplitStrOmitted := string(currentStr[0 : len(currentStr)-splitSliceLen+1])
			strSlice = append(strSlice, strWithSplitStrOmitted)
			currentStr = currentStr[:0]
		} else {
			currentStr = append(currentStr, r[i])
		}
	}
	if len(currentStr) > 0 {
		strSlice = append(strSlice, string(currentStr))
	}
	return strSlice
}

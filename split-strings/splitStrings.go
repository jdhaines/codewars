package splitstrings

func Solution(str string) []string {
	var result []string

	// deal with odd length strings
	if len(str)%2 != 0 {
		str = str + "_"
	}

	counter := 1
	substring := ""

	// Watch if we are on the first or second character of the substring.  On
	// the first, save the char to a substring.  On the second, add it to the
	// substring, then push the substring to the result slice.
	for _, char := range str {
		if counter == 1 { // first substring
			substring += string(char)
			counter = 0
			continue
		} else { // second substring character
			substring += string(char)
			result = append(result, substring)
			substring = ""
			counter = 1
		}
	}

	return result
}

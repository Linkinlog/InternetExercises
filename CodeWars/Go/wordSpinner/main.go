package wordspinner

func main() {

}

func spinWords(s string) string {
	// 0. Create indexOfLastSpace
	indexOfLastSpace := len(s)-1
	// 1. Iterate over the string backwards
	for i := len(s)-1; i >= 0; i-- {
		// 2. If we detect a space drop into processing
		// 3. We want to take the slice from the last time there was a space
		// until now, then replace the original string with what we have
		if string(s[i]) == " " && indexOfLastSpace - i >= 5 {
            //copy(s[2:], s[2:5])
		} else if string(s[i]) == " " {
			indexOfLastSpace = i
		}
	}
	return ""
}

func reverse(letters []string)

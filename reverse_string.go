package reverse_string

func ReverseString(input string) (output string) {
	s := []rune(input)
	for i := len(s) - 1; i >= 0; i-- {
		output += string(s[i])
	}

	return output
}

package iteration

func Repeat(character string, repeatCount int) string {
	var repeatedString string
	for i := 0; i < repeatCount; i++ {
		repeatedString += character
	}
	return repeatedString
}

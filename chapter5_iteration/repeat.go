package iteration

// Repeat will return a provided string the provided number of times
func Repeat(character string, repeats int) string {
	var repeated string
	for i := 0; i < repeats; i++ {
		repeated += character
	}
	return repeated
}

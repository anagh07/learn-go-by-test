package iteration

func Repeat(letter string, n int) string {
	var repeated string
	for i := 0; i < n; i++ {
		repeated += letter
	}
	return repeated
}

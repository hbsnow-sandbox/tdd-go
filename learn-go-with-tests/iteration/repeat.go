package iteration

// Repeat 与えられた文字を指定回繰り返して戻す
func Repeat(character string, count int) string {
	var repeated string
	for i := 0; i < count; i++ {
		repeated += character
	}
	return repeated
}

package iteration

const repeatCount = 5

func Repeat(str string) string {
	var result string

	for i := 0; i < repeatCount; i++ {
		result += str
	}

	return result
}

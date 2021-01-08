package iteration



func Repeat(caracter string, numberOfTimesToRepeat int) string {
	var repeated string
	for i := 0; i < numberOfTimesToRepeat; i++ {
		repeated += caracter
	}
	return repeated
}

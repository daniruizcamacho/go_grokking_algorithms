package recursion

func fibonacci(number int) int {
	if number == 0 || number == 1 {
		return number
	}

	return fibonacci(number-2) + fibonacci(number-1)
}
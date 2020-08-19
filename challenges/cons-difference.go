package challenges

func numsSameConsecDiff(N int, K int) []int {
	numbers := []int{}
	for i := 1; i <= 9; i++ {
		numbers = addDigit(i, N, K, numbers)
	}
	if N == 1 {
		numbers = append(numbers, 0)
	}
	allNumbers := make(map[int]bool)

	setOfNumbers := []int{}

	for _, ele := range numbers {
		if !allNumbers[ele] {
			allNumbers[ele] = true
			setOfNumbers = append(setOfNumbers, ele)
		}
	}
	return setOfNumbers
}
func addDigit(result int, N int, K int, numbers []int) []int {
	if digitsCount(result) == N {
		numbers = append(numbers, result)
		return numbers
	}
	if (result%10)+K <= 9 {
		numbers = addDigit(result*10+((result%10)+K), N, K, numbers)
	}

	if (result%10)-K >= 0 {
		numbers = addDigit(result*10+((result%10)-K), N, K, numbers)
	}

	return numbers
}

func digitsCount(number int) int {
	count := 0
	for number != 0 {
		number /= 10
		count++
	}
	return count

}

package challenges

func sortArrayByParity(A []int) []int {
	var evenIndex, oddIndex int
	n := len(A)
	if n == 0 || n == 1 {
		return A
	}
	evenIndex = getEvenIndex(A)
	oddIndex = getOddIndex(A)
	if evenIndex == -1 || oddIndex == -1 || evenIndex <= oddIndex {
		return A
	}

	for evenIndex < n && oddIndex < n {
		if evenIndex > oddIndex && A[evenIndex]%2 == 0 && A[oddIndex]%2 == 1 {
			A[oddIndex], A[evenIndex] = A[evenIndex], A[oddIndex]
		}
		for evenIndex < n && A[evenIndex]%2 == 1 {
			evenIndex++
		}
		for oddIndex < n && A[oddIndex]%2 == 0 {
			oddIndex++
		}

	}

	return A
}
func getEvenIndex(A []int) int {
	for i := 0; i < len(A); i++ {
		if A[i]%2 == 0 {
			return i
		}
	}
	return -1
}

func getOddIndex(A []int) int {
	for i := 0; i < len(A); i++ {
		if A[i]%2 == 1 {
			return i
		}
	}
	return -1
}

package challenges

import (
	"fmt"
)

func distributeCandies(candies int, numPeople int) []int {
	candiesPeople := make([]int, numPeople)
	roundMultiple := 0
	round := 1
	for candies > 0 {
		for i := 0; i < numPeople; i++ {
			if candies < (i+1)+roundMultiple {
				candiesPeople[i] += candies
				candies = 0
				break
			}
			candiesPeople[i] += (i + 1) + roundMultiple
			candies -= (i + 1) + roundMultiple
		}
		fmt.Printf("After round %d, the distribution is %v", round, candiesPeople)
		roundMultiple = round * numPeople
		fmt.Printf("Round multiples is %d", roundMultiple)
		round++

	}
	return candiesPeople
}

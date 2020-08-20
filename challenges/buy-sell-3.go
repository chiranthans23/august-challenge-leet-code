package challenges

func maxProfit(prices []int) int {
	if len(prices) == 0 || isDesceding(prices) {
		return 0
	}
	if isAsceding(prices) {
		return prices[len(prices)-1] - prices[0]
	}
	var maxProfit int
	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j]-prices[i] > maxProfit {
				maxProfit = prices[j] - prices[i]
			}
		}
	}
	return maxProfit
}

func isAsceding(prices []int) bool {
	ele := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < ele {
			return false
		}
		ele = prices[i]
	}
	return true
}
func isDesceding(prices []int) bool {
	ele := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] > ele {
			return false
		}
		ele = prices[i]
	}
	return true
}

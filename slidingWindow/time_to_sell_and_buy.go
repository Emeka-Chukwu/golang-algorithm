package slidingwindow

// 121. Best Time to Buy and Sell Stock
// Easy
// 28.5K
// 970
// Companies
// You are given an array prices where prices[i] is the price of a given stock on the ith day.

// You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

// Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

// Example 1:

// Input: prices = [7,1,5,3,6,4]
// Output: 5
// Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
// Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.
// Example 2:

// Input: prices = [7,6,4,3,1]
// Output: 0
// Explanation: In this case, no transactions are done and the max profit = 0.

// Constraints:

// 1 <= prices.length <= 105
// 0 <= prices[i] <= 104

func (SlidingWindowAlg) maxProfitee(prices []int) int {
	buy := 0
	// sell := 0
	currentProfit := 0

	for i := 0; i < len(prices)-1; i++ {
		if prices[i+1]-prices[i] > 0 && prices[i+1]-prices[i] > currentProfit {
			currentProfit = prices[i+1] - prices[i]
			if buy == 0 {
				buy = prices[i]
			}
			if prices[i] < buy {
				buy = prices[i]
			}
		} else if prices[i+1]-prices[i] > currentProfit {
			currentProfit = prices[i+1] - prices[i]
			buy = prices[i]
		}
		if prices[i+1]-buy > currentProfit && buy != 0 {
			currentProfit = prices[i+1] - buy
		}
	}
	return currentProfit
}

func (SlidingWindowAlg) maxProfit(prices []int) int {
	l, r := 0, 1
	maxProfit := 0
	for r < len(prices) {
		if prices[l] < prices[r] {
			profit := prices[r] - prices[l]
			if profit > maxProfit {
				maxProfit = profit
			}
		} else {
			l = r
		}
		r++
	}
	return maxProfit
}

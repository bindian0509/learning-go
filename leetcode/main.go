package main

import "fmt"

func main() {

	profit1 := maxProfit([]int{7, 1, 5, 3, 6, 4})
	profit2 := maxProfit([]int{1, 2, 3, 4, 5})
	profit3 := maxProfit([]int{7, 6, 4, 3, 1})
	fmt.Println("--------------------------------")
	fmt.Printf("Profits are %v , %v, %v: \n", profit1, profit2, profit3)
	numodd1 := countOdds(3, 7)
	numodd2 := countOdds(8, 10)
	fmt.Println("--------------------------------")
	fmt.Printf("for 3 and 7 its %v \n", numodd1)
	fmt.Printf("for 8 and 10 its %v \n", numodd2)
	fmt.Println("--------------------------------")

}
func maxProfit(prices []int) int {

	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxProfit += prices[i] - prices[i-1]
		}
	}

	return maxProfit
}

/**
* Runtime: 3 ms, faster than 55.06% of Go online submissions
* Memory Usage: 2 MB, less than 19.03% of Go online submissions
 */
func countOdds(low int, high int) int {

	N := (high - low) / 2
	if high%2 != 0 || low%2 != 0 {
		N++
	}
	return N
}

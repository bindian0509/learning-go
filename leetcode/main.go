package main

import (
	"fmt"
	"math"
)

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

func fib(n int) int {
    goldenRatio := (1 + math.Sqrt(5) ) / 2
	return int(math.Round(math.Pow(goldenRatio, (float64(n)))/math.Sqrt(5)))
}
func fibZeroComplexity(n int) int {
    dp := make([]int, n + 1)
    if n == 0 || n == 1 {
        return n
    }
    dp[0] = 0
    dp[1] = 1
    for i := 2; i <= n; i++ {
        dp[i] = dp[i - 1] + dp[i - 2]
    }
    
    return dp[n]
}
/**
* Runtime: 7 ms, faster than 12.66% 
* Memory Usage: 2.7 MB, less than 89.51%
**/
func runningSum (nums [] int) [] int {

	for i:= 1; i < len (nums) ; i++ {
		nums[i] += nums[i-1]
	}
	return nums 
}
func runningSumFastest(nums []int) []int {
    runningSum := make([]int,len(nums))
    runningSum[0] = nums[0]
    for i:=1; i<len(nums); i++ {
        runningSum[i] = nums[i] + runningSum[i-1]
    }
    return runningSum
}
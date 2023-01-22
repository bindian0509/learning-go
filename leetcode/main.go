package main

import (
	"fmt"
	"math"
)


func is_rotation(arr1 [] int, arr2 [] int) bool {

	if len(arr1) != len(arr2) {
		return false
	}
	loc := -1

	for i := 0; i < len(arr2); i++ {
		if arr2[i] == arr1[0] {
			loc = i
		}
	}
	if loc == -1 {
		return false
	}

	idx := 0

	for idx < len(arr1) {
		if loc == len(arr2) {
			loc = 0
		}

		if arr1[idx] != arr2[loc] {
			return false
		}
		idx++
		loc++
	}


	return true
}


func main() {
	arr1 := [] int {1,2,3,4,5,6,7}
	arr2 := [] int {4, 5, 6, 7, 1, 2, 3, 10, 11}

	fmt.Println(is_rotation(arr1, arr2))
	// false

	array2b := [] int {4, 5, 6, 7, 1, 2, 3}
	fmt.Println(is_rotation(arr1, array2b))
	//# this will print true

	array2c := [] int {4, 5, 6, 9, 1, 2, 3}
	fmt.Println(is_rotation(arr1, array2c))
	//# this will print false

	array2e := [] int {4, 5, 6, 7, 0, 2, 3}
	fmt.Println(is_rotation(arr1, array2e))
	//# this will print false

	array2f := [] int {1, 2, 3, 4, 5, 6, 7}
	fmt.Println(is_rotation(arr1, array2f))
	//# this will print true


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

func maxSubArray(nums []int) int {
	sum := math.MinInt
	curr := 0;

	for i := 0; i < len(nums); i++ {
		curr = curr + nums[i]
		if curr > sum {
			sum = curr
		}
		if curr < 0 {
			curr = 0
		}
	}
	return sum
}

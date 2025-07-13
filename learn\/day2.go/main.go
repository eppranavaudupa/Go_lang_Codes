package main

import "fmt"

func maxProfit(prices []int) int {
	var max int = 0
	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			ele := prices[j] - prices[i]
			if ele > max {
				max = ele
			}
		}
	}
	return max
}
func main() {
	res := maxProfit([]int{7, 1, 5, 3, 6, 4})
	fmt.Println(res)
}

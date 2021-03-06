package main

import (
	"fmt"
)



func main() {
	//股票最大收益
	nums := []int{7, 6, 4, 3, 1}
	fmt.Println(MaxProfit1(nums))
}



//买卖股票的最佳时机
//给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
//
//如果你最多只允许完成一笔交易（即买入和卖出一支股票一次），设计一个算法来计算你所能获取的最大利润。
//
//注意：你不能在买入股票前卖出股票。
//
//
//
//示例 1:
//
//输入: [7,1,5,3,6,4]
//输出: 5
//解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
//注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格。
//示例 2:
//
//输入: [7,6,4,3,1]
//输出: 0
//解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。

//MaxProfit 买卖股票的最佳时机,暴力解法
func MaxProfit(prices []int) int {
	maxVal := 0
	for i := len(prices) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			val := prices[i] - prices[j]
			if val > maxVal {
				maxVal = val
			}
		}
	}
	return maxVal
}

//MaxProfit 买卖股票的最佳时机,剑指offer第63题解法
func MaxProfit1(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	min := prices[0]
	maxDiff := 0
	for i := 1; i < len(prices); i++ {
		if min > prices[i-1] {
			min = prices[i-1]
		}
		currentDiff := prices[i] - min
		if currentDiff > maxDiff {
			maxDiff = prices[i] - min
		}
	}
	return maxDiff
}

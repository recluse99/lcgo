package lcgo

/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 */

// @lc code=start
func maxProfitOne(prices []int) int {

	var Profit int

	//var minPrice int

	//var maxProfit int

	var hasPrice int



	for i := 0; i < len(prices); i++ {

		//首次赋值
		if i==0 {
			hasPrice = prices[i]
			//minPrice=prices[i]
		} 
		if prices[i] >= hasPrice {

			Profit += prices[i] - hasPrice
			//卖出股票，持有为0
			hasPrice = prices[i]

			//当现在持有的价格大于或等于当天的股票价格
		} else {
			hasPrice = prices[i]
			//minPrice = hasPrice
		}

	}

	return Profit
}

// @lc code=end

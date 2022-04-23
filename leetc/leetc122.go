package leetc

func maxProfit122(prices []int) int {
	ret := 0
	buy, sale := prices[0], prices[0]

	for _, v := range prices {
		if v < sale {
			ret += sale - buy
			buy, sale = v, v
			continue
		}
		sale = v
	}
	ret += sale - buy

	return ret
}

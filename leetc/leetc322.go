package leetc

func CoinChange(coins []int, amount int) int {
	count := make([]int, amount+1)

	count[0] = 0
	for i := 1; i < amount+1; i++ {
		mi := amount + 1
		for _, v := range coins {
			if i-v < 0 || count[i-v] == -1 {
				continue
			}
			if count[i-v]+1 < mi {
				mi = count[i-v] + 1
			}
		}
		if mi == amount+1 {
			count[i] = -1
		} else {
			count[i] = mi
		}
	}
	return count[amount]
}

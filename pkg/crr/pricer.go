package crr

func Price_by_CRR(bm BinModel, opt Option) float64 {
	expiry := opt.Expiry()
	prices := make([]float64, expiry+1)
	for i := range prices {
		prices[i] = opt.Payoff(bm.Price(expiry, uint(i)))
	}
	q := bm.MartProb()            // martingale probability
	dc := 1.0 / (1.0 + bm.Rate()) // discount factor
	for j := expiry; j > 0; j-- {
		for i, price := range prices {
			if uint(i) == j {
				break
			}
			prices[i] = dc * (q*price + (1.0-q)*prices[i+1])
		}
	}
	return prices[0]
}

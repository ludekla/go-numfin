package crr

import "math"

// BinModel implements the binomial markt model for one
// underlying. It holds the necessary data for computing the
// price of the underlying, performed by the [BinModel.Price] method.
type BinModel struct {
	spot  float64 // spot price
	utick float64 // uptick
	dtick float64 // downtick
	rate  float64 // risk-neutral rate
	prob  float64 // risk-neutral probability
}

// NewBinModel returns a [BinModel], initialised with the
// passed parameters.
func NewBinModel(spot float64, ut float64, dt float64, rate float64) BinModel {
	q := (rate - dt) / (ut - dt)
	return BinModel{spot, ut, dt, rate, q}
}

// BinModel.Price computes the price of the underlying for the given
// expiry and number of upticks.
func (b BinModel) Price(expiry uint, nUticks uint) float64 {
	lu := float64(nUticks) * math.Log(1.0+b.utick)
	ld := float64(expiry-nUticks) * math.Log(1.0+b.dtick)
	return b.spot * math.Exp(lu+ld)
}

func (b BinModel) Rate() float64 {
	return b.rate
}

func (b BinModel) MartProb() float64 {
	return b.prob
}

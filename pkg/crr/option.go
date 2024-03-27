// Package crr implements the CRR pricing model.
package crr

// Option is an interface for the CRR pricer. [Expiry] is meant
// to be a getter method, [Payoff] returns the option's payoff
// at the given price of the underlying.
type Option interface {
	Expiry() uint
	Payoff(price float64) float64
}

// EuroOption holds the expiry of a European option
// and implements the [Expiry] method. It serves as
// an object to be embedded.
type EurOption struct {
	expiry uint
}

// CallOption implements the [Option] interface and represents
// a call option in the CRR model with expiry and strike price.
type CallOption struct {
	EurOption
	strike float64
}

// PutOption
type PutOption struct {
	EurOption
	strike float64
}

type DigitCall struct {
	EurOption
	strike float64
}

type DigitPut struct {
	EurOption
	strike float64
}

func (v EurOption) Expiry() uint {
	return v.expiry
}

// NewCallOption returns a [CallOption] initialised by the
// two arguments expiry for the embedded [EurOption] and
// the strike price.
func NewCallOption(expiry uint, strike float64) CallOption {
	return CallOption{EurOption{expiry}, strike}
}

// Payoff implements the [Option.Payoff] method which makes
// the CallOption conform to the [Option] interface.
func (c CallOption) Payoff(price float64) float64 {
	if price > c.strike {
		return price - c.strike
	}
	return 0.0
}

func NewPutOption(expiry uint, strike float64) PutOption {
	return PutOption{EurOption{expiry}, strike}
}

func (p PutOption) Payoff(price float64) float64 {
	if price < p.strike {
		return p.strike - price
	}
	return 0.0
}

func NewDigitCall(expiry uint, strike float64) DigitCall {
	return DigitCall{EurOption{expiry}, strike}
}

func (c DigitCall) Payoff(price float64) float64 {
	if price > c.strike {
		return 1.0
	}
	return 0.0
}

func NewDigitPut(expiry uint, strike float64) DigitPut {
	return DigitPut{EurOption{expiry}, strike}
}

func (p DigitPut) Payoff(price float64) float64 {
	if price < p.strike {
		return 1.0
	}
	return 0.0
}

// Package crr implements the CRR pricing model.
package crr

// Option is an interface for the CRR pricer. [Expiry] is meant
// to be a getter method for the expiry field, [Payoff] returns
// the option's payoff at the given price of the underlying.
type Option interface {
	Expiry() int
	Payoff(price float64) float64
}

// EuroOption holds the expiry of a European option
// and implements the [Expiry] method. It serves as
// an object to be embedded.
type EurOption struct {
	expiry int
}

func (v EurOption) Expiry() int {
	return v.expiry
}

// CallOption implements the [Option] interface for a call
// option in the CRR model with expiry and strike price.
type CallOption struct {
	EurOption
	strike float64
}

// NewCallOption returns a [CallOption] initialised by the
// two arguments expiry for the embedded [EurOption] and
// the strike price.
func NewCallOption(expiry int, strike float64) CallOption {
	return CallOption{EurOption{expiry}, strike}
}

func (c CallOption) Payoff(price float64) float64 {
	if price > c.strike {
		return price - c.strike
	}
	return 0.0
}

// PutOption implements the [Option] interface for a put
// option in the CRR model with expiry and strike price.
type PutOption struct {
	EurOption
	strike float64
}

// NewPutOption returns a [PutOption] with the given expiry
// and strike price.
func NewPutOption(expiry int, strike float64) PutOption {
	return PutOption{EurOption{expiry}, strike}
}

func (p PutOption) Payoff(price float64) float64 {
	if price < p.strike {
		return p.strike - price
	}
	return 0.0
}

// DigitCall implements the [Option] interface for a digital
// call option in the CRR model with expiry and strike price.
type DigitCall struct {
	EurOption
	strike float64
}

// NewDigitalCall returns a [DigitCall] with expiry and
// strike as given.
func NewDigitCall(expiry int, strike float64) DigitCall {
	return DigitCall{EurOption{expiry}, strike}
}

func (c DigitCall) Payoff(price float64) float64 {
	if price > c.strike {
		return 1.0
	}
	return 0.0
}

// DigitPut implements the [Option] interface for a digital
// put option in the CRR model with expiry and strike price.
type DigitPut struct {
	EurOption
	strike float64
}

// NewDigitPut returns a [DigitPut] with expiry and
// strike as given.
func NewDigitPut(expiry int, strike float64) DigitPut {
	return DigitPut{EurOption{expiry}, strike}
}

func (p DigitPut) Payoff(price float64) float64 {
	if price < p.strike {
		return 1.0
	}
	return 0.0
}

// DoubleDigit implements the [Option] interface for a
// double digital option.
type DoubleDigit struct {
	EurOption
	strikeLo float64
	strikeHi float64
}

// NewDoubleDigit return a [DoubleDigit] object initialised with the
// given parameters.
func NewDoubleDigit(expiry int, lo float64, hi float64) DoubleDigit {
	return DoubleDigit{EurOption{expiry}, lo, hi}
}

func (d DoubleDigit) Payoff(price float64) float64 {
	if price >= d.strikeLo && price <= d.strikeHi {
		return 1.0
	}
	return 0.0
}

// BearSpread implements the [Option] interface for a
// bear spread option.
type BearSpread struct {
	EurOption
	strikeLo float64
	strikeHi float64
}

// NewBearSpread return a [BearSpread] object initialised with the
// given parameters.
func NewBearSpread(expiry int, lo float64, hi float64) BearSpread {
	return BearSpread{EurOption{expiry}, lo, hi}
}

func (b BearSpread) Payoff(price float64) float64 {
	if price < b.strikeLo {
		return b.strikeHi - b.strikeLo
	} else if price > b.strikeHi {
		return 0.0
	}
	return b.strikeHi - price
}

// BullSpread implements the [Option] interface for a
// bear spread option.
type BullSpread struct {
	EurOption
	strikeLo float64
	strikeHi float64
}

// NewBullSpread return a [BullSpread] object initialised with the
// given parameters.
func NewBullSpread(expiry int, lo float64, hi float64) BullSpread {
	return BullSpread{EurOption{expiry}, lo, hi}
}

func (b BullSpread) Payoff(price float64) float64 {
	if price > b.strikeHi {
		return b.strikeHi - b.strikeLo
	} else if price < b.strikeLo {
		return 0.0
	}
	return price - b.strikeLo
}

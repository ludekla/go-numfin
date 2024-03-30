package main

import (
	"fmt"
	"numfin/pkg/crr"
)

func main() {
	fmt.Println("Hello CRR Option Pricer!")

	// Market Data
	const (
		spot = 100.0
		up   = 0.01
		down = -0.01
		rate = 0.005
	)
	// Set up market model.
	var bm = crr.NewBinModel(spot, down, up, rate)

	// Option Data
	const (
		expiry  = 100
		cstrike = 100.0
		pstrike = 200.0
	)
	// Set up options.
	var (
		// plain vanilla
		call  = crr.NewCallOption(expiry, cstrike)
		put   = crr.NewPutOption(expiry, pstrike)
		dcall = crr.NewDigitCall(expiry, cstrike)
		dput  = crr.NewDigitPut(expiry, pstrike)
		// spreads
		dopt = crr.NewDoubleDigit(expiry, cstrike, pstrike)
		bull = crr.NewBullSpread(expiry, cstrike, pstrike)
		bear = crr.NewBearSpread(expiry, cstrike, pstrike)
	)

	// Let fly.
	// plain vanilla
	cp := crr.Price_by_CRR(bm, call)
	pp := crr.Price_by_CRR(bm, put)
	fmt.Printf("Call:         %.5f Put:         %.5f\n", cp, pp)
	dc := crr.Price_by_CRR(bm, dcall)
	dp := crr.Price_by_CRR(bm, dput)
	fmt.Printf("Digital Call: %.5f  Digital Put: %.5f\n", dc, dp)
	// Spreads
	blp := crr.Price_by_CRR(bm, bull)
	brp := crr.Price_by_CRR(bm, bear)
	fmt.Printf("Bull:         %.5f Bear:        %.5f\n", blp, brp)
	ddp := crr.Price_by_CRR(bm, dopt)
	fmt.Printf("Double-Digit Option: %.5f\n", ddp)

}

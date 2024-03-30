package crr

import (
	"math"
	"testing"
)

const (
	spot   = 100.0
	up     = 0.01
	down   = -0.01
	rate   = 0.005
	expiry = 100
)

func TestCallPrices(t *testing.T) {
	bm := NewBinModel(spot, down, up, rate)
	strike := spot
	call := NewCallOption(expiry, strike)
	if price := 39.271; math.Abs(Price_by_CRR(bm, call)-price) > 1e-3 {
		t.Errorf("expected crr_price: %v, got %v", price, Price_by_CRR(bm, call))
	}
	dcall := NewDigitCall(expiry, strike)
	if price := 0.607; math.Abs(Price_by_CRR(bm, dcall)-price) > 1e-3 {
		t.Errorf("expected crr_price: %v, got %v", price, Price_by_CRR(bm, dcall))
	}
}

func TestPutPrices(t *testing.T) {
	bm := NewBinModel(spot, up, down, rate)
	strike := 2.0 * spot
	put := NewPutOption(expiry, strike)
	if price := 21.484; math.Abs(Price_by_CRR(bm, put)-price) > 1e-3 {
		t.Errorf("expected crr_price: %v, got %v", price, Price_by_CRR(bm, put))
	}
	dput := NewDigitPut(expiry, strike)
	if price := 0.601; math.Abs(Price_by_CRR(bm, dput)-price) > 1e-3 {
		t.Errorf("expected crr_price: %v, got %v", price, Price_by_CRR(bm, dput))
	}
}

func TestSpreadPrices(t *testing.T) {
	bm := NewBinModel(spot, up, down, rate)
	his, los := 2*spot, spot // high strike, low strike
	dopt := NewDoubleDigit(expiry, los, his)
	if price := 0.601; math.Abs(Price_by_CRR(bm, dopt)-price) > 1e-3 {
		t.Errorf("expected crr_price: %v, got %v", price, Price_by_CRR(bm, dopt))
	}
	bull := NewBullSpread(expiry, los, his)
	if price := 39.245; math.Abs(Price_by_CRR(bm, bull)-price) > 1e-3 {
		t.Errorf("expected crr_price: %v, got %v", price, Price_by_CRR(bm, bull))
	}
	bear := NewBearSpread(expiry, los, his)
	if price := 21.484; math.Abs(Price_by_CRR(bm, bear)-price) > 1e-3 {
		t.Errorf("expected crr_price: %v, got %v", price, Price_by_CRR(bm, bear))
	}
}

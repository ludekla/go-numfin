package crr

import (
	"math"
	"testing"
)

func TestPrice(t *testing.T) {
	bm := NewBinModel(100.0, 0.1, -0.1, 0.02)
	dcall := NewDigitCall(1, 100.0)
	if price := 101.2; math.Abs(Price_by_CRR(bm, dcall)-price) > 1e-4 {
		t.Errorf("expected crr_price: %v, got %v", price, Price_by_CRR(bm, dcall))
	}
}

package crr

import (
	"math"
	"testing"
)

func TestBinModelPrice(t *testing.T) {
	bm := NewBinModel(100.0, 0.1, -0.1)
	if price := 450.05074718778; math.Abs(price-bm.Price(100, 60)) > 1e-3 {
		t.Errorf("expected Price() to be %v, got %v", price, bm.Price(100, 60))
	}
}

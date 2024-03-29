package crr

import (
	"math"
	"testing"
)

func TestCallOption(t *testing.T) {
	const N uint = 100
	call := NewCallOption(N, 103.1)
	if call.Expiry() != N {
		t.Errorf("expected Expiry() == %d got %v", N, call.Expiry())
	}
	poff := 1.1
	if price := 104.2; math.Abs(call.Payoff(price)-poff) > 1e-10 {
		t.Errorf("expected Payoff() == %f got %f", poff, call.Payoff(price))
	}
	poff = 0.0
	if price := 100.0; math.Abs(call.Payoff(price)-poff) > 1e-10 {
		t.Errorf("expected Payoff() == %f got %f", poff, call.Payoff(price))
	}
}

func TestPutOption(t *testing.T) {
	const N uint = 100
	put := NewPutOption(N, 103.1)
	if put.Expiry() != N {
		t.Errorf("expected Expiry() == %d got %v", N, put.Expiry())
	}
	poff := 0.0
	if price := 104.2; math.Abs(put.Payoff(price)-poff) > 1e-10 {
		t.Errorf("expected Payoff() == %f got %f", poff, put.Payoff(price))
	}
	poff = 0.2
	if price := 102.9; math.Abs(put.Payoff(price)-poff) > 1e-10 {
		t.Errorf("expected Payoff() == %f got %f", poff, put.Payoff(price))
	}
}

func TestDigitCall(t *testing.T) {
	const N uint = 100
	dcall := NewDigitCall(N, 103.1)
	if dcall.Expiry() != N {
		t.Errorf("expected Expiry() == %d got %v", N, dcall.Expiry())
	}
	poff := 1.0
	if price := 104.2; math.Abs(dcall.Payoff(price)-poff) > 1e-10 {
		t.Errorf("expected Payoff() == %f got %f", poff, dcall.Payoff(price))
	}
	poff = 0.0
	if price := 100.0; math.Abs(dcall.Payoff(price)-poff) > 1e-10 {
		t.Errorf("expected Payoff() == %f got %f", poff, dcall.Payoff(price))
	}
}

func TestDigitPut(t *testing.T) {
	const N uint = 100
	dput := NewDigitPut(N, 103.1)
	if dput.Expiry() != N {
		t.Errorf("expected Expiry() == %d got %v", N, dput.Expiry())
	}
	poff := 0.0
	if price := 104.2; math.Abs(dput.Payoff(price)-poff) > 1e-10 {
		t.Errorf("expected Payoff() == %f got %f", poff, dput.Payoff(price))
	}
	poff = 1.0
	if price := 102.9; math.Abs(dput.Payoff(price)-poff) > 1e-10 {
		t.Errorf("expected Payoff() == %f got %f", poff, dput.Payoff(price))
	}
}

func TestDoubleDigit(t *testing.T) {
	const N uint = 100
	dod := NewDoubleDigit(N, 98.2, 103.1)
	if dod.Expiry() != N {
		t.Errorf("expected Expiry() == %d got %v", N, dod.Expiry())
	}
	poff := 0.0
	if price := 104.2; math.Abs(dod.Payoff(price)-poff) > 1e-10 {
		t.Errorf("expected Payoff() == %f got %f", poff, dod.Payoff(price))
	}
	if price := 98.0; math.Abs(dod.Payoff(price)-poff) > 1e-10 {
		t.Errorf("expected Payoff() == %f got %f", poff, dod.Payoff(price))
	}
	poff = 1.0
	if price := 102.9; math.Abs(dod.Payoff(price)-poff) > 1e-10 {
		t.Errorf("expected Payoff() == %f got %f", poff, dod.Payoff(price))
	}
}

func TestBearSpread(t *testing.T) {
	const N uint = 100
	bear := NewBearSpread(N, 98.2, 103.1)
	if bear.Expiry() != N {
		t.Errorf("expected Expiry() == %d got %v", N, bear.Expiry())
	}
	poff := 0.0
	if price := 104.2; math.Abs(bear.Payoff(price)-poff) > 1e-10 {
		t.Errorf("expected Payoff() == %f got %f", poff, bear.Payoff(price))
	}
	poff = 4.9
	if price := 98.0; math.Abs(bear.Payoff(price)-poff) > 1e-10 {
		t.Errorf("expected Payoff() == %f got %f", poff, bear.Payoff(price))
	}
	poff = 0.2
	if price := 102.9; math.Abs(bear.Payoff(price)-poff) > 1e-10 {
		t.Errorf("expected Payoff() == %f got %f", poff, bear.Payoff(price))
	}
}

func TestBullSpread(t *testing.T) {
	const N uint = 100
	bull := NewBullSpread(N, 98.2, 103.1)
	if bull.Expiry() != N {
		t.Errorf("expected Expiry() == %d got %v", N, bull.Expiry())
	}
	poff := 4.9
	if price := 104.2; math.Abs(bull.Payoff(price)-poff) > 1e-10 {
		t.Errorf("expected Payoff() == %f got %f", poff, bull.Payoff(price))
	}
	poff = 0.0
	if price := 98.0; math.Abs(bull.Payoff(price)-poff) > 1e-10 {
		t.Errorf("expected Payoff() == %f got %f", poff, bull.Payoff(price))
	}
	poff = 4.7
	if price := 102.9; math.Abs(bull.Payoff(price)-poff) > 1e-10 {
		t.Errorf("expected Payoff() == %f got %f", poff, bull.Payoff(price))
	}
}

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
	if price := 104.2; math.Abs(call.Payoff(price)-poff) > 1e-5 {
		t.Errorf("expected Payoff() == %f got %f", poff, call.Payoff(price))
	}
}

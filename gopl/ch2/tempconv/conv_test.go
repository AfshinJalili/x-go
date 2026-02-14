package tempconv

import "testing"

func TestCToF(t *testing.T) {
	if got := CToF(FreezingC); got != 32 {
		t.Errorf("CToF(0) = %v, want 32", got)
	}
	if got := CToF(BoilingC); got != 212 {
		t.Errorf("CToF(100) = %v, want 212", got)
	}
}

func TestFToC(t *testing.T) {
	if got := FToC(32); got != FreezingC {
		t.Errorf("FToC(32) = %v, want 0", got)
	}
	if got := FToC(212); got != BoilingC {
		t.Errorf("FToC(212) = %v, want 100", got)
	}
}

func TestKToC(t *testing.T) {
	if got := KToC(FreezingK); got != FreezingC {
		t.Errorf("KToC(273.15) = %v, want 0", got)
	}
	if got := KToC(BoilingK); got != BoilingC {
		t.Errorf("KToC(373.15) = %v, want 100", got)
	}
}

func TestCToK(t *testing.T) {
	if got := CToK(FreezingC); got != FreezingK {
		t.Errorf("CToK(0) = %v, want 273.15", got)
	}
	if got := CToK(BoilingC); got != BoilingK {
		t.Errorf("CToK(100) = %v, want 373.15", got)
	}
}

func TestKToF(t *testing.T) {
	if got := KToF(FreezingK); got != 32 {
		t.Errorf("KToF(273.15) = %v, want 32", got)
	}
}

func TestFToK(t *testing.T) {
	if got := FToK(32); got != FreezingK {
		t.Errorf("FToK(32) = %v, want 273.15", got)
	}
}

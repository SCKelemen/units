package units

import "testing"

func TestResolutionConstructors(t *testing.T) {
	tests := []struct {
		name     string
		res      Resolution
		expected string
		unit     ResolutionUnit
	}{
		{"Dpi", Dpi(96), "96.00dpi", DotsPerInch},
		{"Dpcm", Dpcm(37.795), "37.80dpcm", DotsPerCentimeter},
		{"Dppx", Dppx(2), "2.00dppx", DotsPerPixel},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.res.Unit != tt.unit {
				t.Errorf("Unit = %s, want %s", tt.res.Unit, tt.unit)
			}
			if tt.res.String() != tt.expected {
				t.Errorf("String() = %s, want %s", tt.res.String(), tt.expected)
			}
		})
	}
}

func TestResolutionConversions(t *testing.T) {
	tests := []struct {
		name     string
		res      Resolution
		expected Resolution
	}{
		{"96dpi to dppx", Dpi(96).ToDppx(), Dppx(1)},
		{"192dpi to dppx", Dpi(192).ToDppx(), Dppx(2)},
		{"1dppx to dpi", Dppx(1).ToDpi(), Dpi(96)},
		{"2dppx to dpi", Dppx(2).ToDpi(), Dpi(192)},
		{"96dpi to dpcm", Dpi(96).ToDpcm(), Dpcm(37.795275591)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !almostEqual(tt.res.Value, tt.expected.Value) {
				t.Errorf("Value = %.6f%s, want %.6f%s",
					tt.res.Value, tt.res.Unit, tt.expected.Value, tt.expected.Unit)
			}
			if tt.res.Unit != tt.expected.Unit {
				t.Errorf("Unit = %s, want %s", tt.res.Unit, tt.expected.Unit)
			}
		})
	}
}

func TestResolutionArithmetic(t *testing.T) {
	t.Run("Add", func(t *testing.T) {
		result := Dpi(96).Add(Dpi(96))
		if result.Value != 192.0 {
			t.Errorf("Add() = %.2f, want 192.00", result.Value)
		}
	})

	t.Run("Sub", func(t *testing.T) {
		result := Dpi(192).Sub(Dpi(96))
		if result.Value != 96.0 {
			t.Errorf("Sub() = %.2f, want 96.00", result.Value)
		}
	})

	t.Run("Mul", func(t *testing.T) {
		result := Dppx(1).Mul(2.0)
		if result.Value != 2.0 {
			t.Errorf("Mul() = %.2f, want 2.00", result.Value)
		}
	})

	t.Run("Div", func(t *testing.T) {
		result := Dppx(4).Div(2.0)
		if result.Value != 2.0 {
			t.Errorf("Div() = %.2f, want 2.00", result.Value)
		}
	})
}

func TestResolutionComparison(t *testing.T) {
	tests := []struct {
		name        string
		a           Resolution
		b           Resolution
		lessThan    bool
		greaterThan bool
	}{
		{"96dpi < 192dpi", Dpi(96), Dpi(192), true, false},
		{"192dpi > 96dpi", Dpi(192), Dpi(96), false, true},
		{"1dppx < 2dppx", Dppx(1), Dppx(2), true, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.a.LessThan(tt.b) != tt.lessThan {
				t.Errorf("LessThan() = %v, want %v", tt.a.LessThan(tt.b), tt.lessThan)
			}
			if tt.a.GreaterThan(tt.b) != tt.greaterThan {
				t.Errorf("GreaterThan() = %v, want %v", tt.a.GreaterThan(tt.b), tt.greaterThan)
			}
		})
	}
}

func TestResolutionIsZero(t *testing.T) {
	tests := []struct {
		name     string
		res      Resolution
		expected bool
	}{
		{"Zero dpi", Dpi(0), true},
		{"Non-zero dpi", Dpi(96), false},
		{"Zero dppx", Dppx(0), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.res.IsZero() != tt.expected {
				t.Errorf("IsZero() = %v, want %v", tt.res.IsZero(), tt.expected)
			}
		})
	}
}

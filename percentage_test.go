package units

import "testing"

func TestPercentageConstructor(t *testing.T) {
	tests := []struct {
		name     string
		pct      Percentage
		expected string
		value    float64
	}{
		{"50%", Percent(50), "50.00%", 50.0},
		{"100%", Percent(100), "100.00%", 100.0},
		{"0%", Percent(0), "0.00%", 0.0},
		{"25.5%", Percent(25.5), "25.50%", 25.5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.pct.Value != tt.value {
				t.Errorf("Value = %f, want %f", tt.pct.Value, tt.value)
			}
			if tt.pct.String() != tt.expected {
				t.Errorf("String() = %s, want %s", tt.pct.String(), tt.expected)
			}
		})
	}
}

func TestPercentageFraction(t *testing.T) {
	tests := []struct {
		name     string
		pct      Percentage
		expected float64
	}{
		{"50% = 0.5", Percent(50), 0.5},
		{"100% = 1.0", Percent(100), 1.0},
		{"25% = 0.25", Percent(25), 0.25},
		{"0% = 0.0", Percent(0), 0.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.pct.Fraction()
			if !almostEqual(result, tt.expected) {
				t.Errorf("Fraction() = %f, want %f", result, tt.expected)
			}
		})
	}
}

func TestPercentageOf(t *testing.T) {
	tests := []struct {
		name     string
		pct      Percentage
		value    float64
		expected float64
	}{
		{"50% of 200", Percent(50), 200, 100},
		{"25% of 100", Percent(25), 100, 25},
		{"100% of 50", Percent(100), 50, 50},
		{"10% of 1000", Percent(10), 1000, 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.pct.Of(tt.value)
			if !almostEqual(result, tt.expected) {
				t.Errorf("Of(%f) = %f, want %f", tt.value, result, tt.expected)
			}
		})
	}
}

func TestPercentageArithmetic(t *testing.T) {
	t.Run("Add", func(t *testing.T) {
		result := Percent(50).Add(Percent(25))
		if result.Value != 75.0 {
			t.Errorf("Add() = %f, want 75.0", result.Value)
		}
	})

	t.Run("Sub", func(t *testing.T) {
		result := Percent(100).Sub(Percent(30))
		if result.Value != 70.0 {
			t.Errorf("Sub() = %f, want 70.0", result.Value)
		}
	})

	t.Run("Mul", func(t *testing.T) {
		result := Percent(50).Mul(2.0)
		if result.Value != 100.0 {
			t.Errorf("Mul() = %f, want 100.0", result.Value)
		}
	})

	t.Run("Div", func(t *testing.T) {
		result := Percent(100).Div(2.0)
		if result.Value != 50.0 {
			t.Errorf("Div() = %f, want 50.0", result.Value)
		}
	})
}

func TestPercentageComparison(t *testing.T) {
	tests := []struct {
		name        string
		a           Percentage
		b           Percentage
		lessThan    bool
		greaterThan bool
		equals      bool
	}{
		{"25% < 50%", Percent(25), Percent(50), true, false, false},
		{"75% > 50%", Percent(75), Percent(50), false, true, false},
		{"50% = 50%", Percent(50), Percent(50), false, false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.a.LessThan(tt.b) != tt.lessThan {
				t.Errorf("LessThan() = %v, want %v", tt.a.LessThan(tt.b), tt.lessThan)
			}
			if tt.a.GreaterThan(tt.b) != tt.greaterThan {
				t.Errorf("GreaterThan() = %v, want %v", tt.a.GreaterThan(tt.b), tt.greaterThan)
			}
			if tt.a.Equals(tt.b) != tt.equals {
				t.Errorf("Equals() = %v, want %v", tt.a.Equals(tt.b), tt.equals)
			}
		})
	}
}

func TestPercentageClamp(t *testing.T) {
	tests := []struct {
		name     string
		pct      Percentage
		min      float64
		max      float64
		expected float64
	}{
		{"Clamp 150% to 0-100", Percent(150), 0, 100, 100},
		{"Clamp -10% to 0-100", Percent(-10), 0, 100, 0},
		{"Clamp 50% to 0-100", Percent(50), 0, 100, 50},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.pct.Clamp(tt.min, tt.max)
			if result.Value != tt.expected {
				t.Errorf("Clamp(%f, %f) = %f, want %f", tt.min, tt.max, result.Value, tt.expected)
			}
		})
	}
}

func TestPercentageIsZero(t *testing.T) {
	tests := []struct {
		name     string
		pct      Percentage
		expected bool
	}{
		{"0% is zero", Percent(0), true},
		{"50% is not zero", Percent(50), false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.pct.IsZero() != tt.expected {
				t.Errorf("IsZero() = %v, want %v", tt.pct.IsZero(), tt.expected)
			}
		})
	}
}

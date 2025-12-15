package units

import "testing"

func TestFrequencyConstructors(t *testing.T) {
	tests := []struct {
		name     string
		freq     Frequency
		expected string
		unit     FrequencyUnit
	}{
		{"Hertz", Hz(440), "440.00Hz", Hertz},
		{"Kilohertz", KHz(20), "20.00kHz", Kilohertz},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.freq.Unit != tt.unit {
				t.Errorf("Unit = %s, want %s", tt.freq.Unit, tt.unit)
			}
			if tt.freq.String() != tt.expected {
				t.Errorf("String() = %s, want %s", tt.freq.String(), tt.expected)
			}
		})
	}
}

func TestFrequencyConversions(t *testing.T) {
	tests := []struct {
		name     string
		freq     Frequency
		expected Frequency
	}{
		{"1kHz to Hz", KHz(1).ToHz(), Hz(1000)},
		{"2.5kHz to Hz", KHz(2.5).ToHz(), Hz(2500)},
		{"1000Hz to kHz", Hz(1000).ToKHz(), KHz(1)},
		{"500Hz to kHz", Hz(500).ToKHz(), KHz(0.5)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !almostEqual(tt.freq.Value, tt.expected.Value) {
				t.Errorf("Value = %.6f%s, want %.6f%s",
					tt.freq.Value, tt.freq.Unit, tt.expected.Value, tt.expected.Unit)
			}
			if tt.freq.Unit != tt.expected.Unit {
				t.Errorf("Unit = %s, want %s", tt.freq.Unit, tt.expected.Unit)
			}
		})
	}
}

func TestFrequencyArithmetic(t *testing.T) {
	t.Run("Add", func(t *testing.T) {
		result := Hz(440).Add(Hz(220))
		if result.Value != 660.0 {
			t.Errorf("Add() = %.2f, want 660.00", result.Value)
		}
	})

	t.Run("Sub", func(t *testing.T) {
		result := Hz(1000).Sub(Hz(500))
		if result.Value != 500.0 {
			t.Errorf("Sub() = %.2f, want 500.00", result.Value)
		}
	})

	t.Run("Mul", func(t *testing.T) {
		result := KHz(1.5).Mul(2.0)
		if result.Value != 3.0 {
			t.Errorf("Mul() = %.2f, want 3.00", result.Value)
		}
	})

	t.Run("Div", func(t *testing.T) {
		result := KHz(3).Div(2.0)
		if result.Value != 1.5 {
			t.Errorf("Div() = %.2f, want 1.50", result.Value)
		}
	})
}

func TestFrequencyComparison(t *testing.T) {
	tests := []struct {
		name        string
		a           Frequency
		b           Frequency
		lessThan    bool
		greaterThan bool
	}{
		{"440Hz < 880Hz", Hz(440), Hz(880), true, false},
		{"880Hz > 440Hz", Hz(880), Hz(440), false, true},
		{"1kHz < 2kHz", KHz(1), KHz(2), true, false},
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

func TestFrequencyIsZero(t *testing.T) {
	tests := []struct {
		name     string
		freq     Frequency
		expected bool
	}{
		{"Zero hertz", Hz(0), true},
		{"Non-zero hertz", Hz(440), false},
		{"Zero kilohertz", KHz(0), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.freq.IsZero() != tt.expected {
				t.Errorf("IsZero() = %v, want %v", tt.freq.IsZero(), tt.expected)
			}
		})
	}
}

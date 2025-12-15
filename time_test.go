package units

import "testing"

func TestTimeConstructors(t *testing.T) {
	tests := []struct {
		name     string
		time     Time
		expected string
		unit     TimeUnit
	}{
		{"Seconds", Sec(2.5), "2.50s", Second},
		{"Milliseconds", Ms(500), "500.00ms", Millisecond},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.time.Unit != tt.unit {
				t.Errorf("Unit = %s, want %s", tt.time.Unit, tt.unit)
			}
			if tt.time.String() != tt.expected {
				t.Errorf("String() = %s, want %s", tt.time.String(), tt.expected)
			}
		})
	}
}

func TestTimeConversions(t *testing.T) {
	tests := []struct {
		name     string
		time     Time
		expected Time
	}{
		{"1s to ms", Sec(1).ToMs(), Ms(1000)},
		{"2.5s to ms", Sec(2.5).ToMs(), Ms(2500)},
		{"1000ms to s", Ms(1000).ToSec(), Sec(1)},
		{"500ms to s", Ms(500).ToSec(), Sec(0.5)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !almostEqual(tt.time.Value, tt.expected.Value) {
				t.Errorf("Value = %.6f%s, want %.6f%s",
					tt.time.Value, tt.time.Unit, tt.expected.Value, tt.expected.Unit)
			}
			if tt.time.Unit != tt.expected.Unit {
				t.Errorf("Unit = %s, want %s", tt.time.Unit, tt.expected.Unit)
			}
		})
	}
}

func TestTimeArithmetic(t *testing.T) {
	t.Run("Add", func(t *testing.T) {
		result := Sec(1).Add(Sec(0.5))
		if result.Value != 1.5 {
			t.Errorf("Add() = %.2f, want 1.50", result.Value)
		}
	})

	t.Run("Sub", func(t *testing.T) {
		result := Sec(2).Sub(Sec(0.5))
		if result.Value != 1.5 {
			t.Errorf("Sub() = %.2f, want 1.50", result.Value)
		}
	})

	t.Run("Mul", func(t *testing.T) {
		result := Sec(1.5).Mul(2.0)
		if result.Value != 3.0 {
			t.Errorf("Mul() = %.2f, want 3.00", result.Value)
		}
	})

	t.Run("Div", func(t *testing.T) {
		result := Sec(3).Div(2.0)
		if result.Value != 1.5 {
			t.Errorf("Div() = %.2f, want 1.50", result.Value)
		}
	})
}

func TestTimeComparison(t *testing.T) {
	tests := []struct {
		name        string
		a           Time
		b           Time
		lessThan    bool
		greaterThan bool
	}{
		{"1s < 2s", Sec(1), Sec(2), true, false},
		{"2s > 1s", Sec(2), Sec(1), false, true},
		{"500ms < 1000ms", Ms(500), Ms(1000), true, false},
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

func TestTimeIsZero(t *testing.T) {
	tests := []struct {
		name     string
		time     Time
		expected bool
	}{
		{"Zero seconds", Sec(0), true},
		{"Non-zero seconds", Sec(1), false},
		{"Zero milliseconds", Ms(0), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.time.IsZero() != tt.expected {
				t.Errorf("IsZero() = %v, want %v", tt.time.IsZero(), tt.expected)
			}
		})
	}
}

package units

import (
	"math"
	"testing"
)

const epsilon = 0.0001 // For floating point comparisons

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) < epsilon
}

func TestAbsoluteLengthConversions(t *testing.T) {
	tests := []struct {
		name     string
		from     Length
		expected Length
	}{
		// Identity conversions
		{"100px to px", Px(100), Px(100)},

		// Inch conversions (1in = 96px)
		{"1in to px", In(1), Px(96)},

		// Centimeter conversions (1in = 2.54cm = 96px)
		{"1cm to px", Cm(1), Px(37.795275591)},
		{"2.54cm to px", Cm(2.54), Px(96)},

		// Point conversions (72pt = 1in = 96px)
		{"72pt to px", Pt(72), Px(96)},
		{"1pt to px", Pt(1), Px(96.0 / 72.0)},

		// Pica conversions (1pc = 12pt = 1/6in)
		{"1pc to px", Pc(1), Px(16)},
		{"6pc to px", Pc(6), Px(96)},

		// Millimeter conversions (10mm = 1cm)
		{"10mm to px", Mm(10), Px(37.795275591)},
		{"25.4mm to px", Mm(25.4), Px(96)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.from.ToPx()
			if err != nil {
				t.Fatalf("ToPx() error = %v", err)
			}
			if !almostEqual(result.Value, tt.expected.Value) {
				t.Errorf("ToPx() = %.6fpx, want %.6fpx", result.Value, tt.expected.Value)
			}
			if result.Unit != PX {
				t.Errorf("ToPx() unit = %s, want %s", result.Unit, PX)
			}
		})
	}
}

func TestAbsoluteLengthTo(t *testing.T) {
	tests := []struct {
		name       string
		from       Length
		targetUnit LengthUnit
		expected   float64
	}{
		// Pixel to other units
		{"96px to in", Px(96), IN, 1.0},
		{"96px to pt", Px(96), PT, 72.0},
		{"96px to pc", Px(96), PC, 6.0},

		// Inch to other units
		{"1in to cm", In(1), CM, 2.54},
		{"1in to mm", In(1), MM, 25.4},
		{"1in to pt", In(1), PT, 72.0},

		// Cm to other units
		{"2.54cm to in", Cm(2.54), IN, 1.0},
		{"1cm to mm", Cm(1), MM, 10.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.from.To(tt.targetUnit)
			if err != nil {
				t.Fatalf("To() error = %v", err)
			}
			if !almostEqual(result.Value, tt.expected) {
				t.Errorf("To(%s) = %.6f%s, want %.6f%s",
					tt.targetUnit, result.Value, result.Unit, tt.expected, tt.targetUnit)
			}
			if result.Unit != tt.targetUnit {
				t.Errorf("To() unit = %s, want %s", result.Unit, tt.targetUnit)
			}
		})
	}
}

func TestToPxErrorOnRelativeUnits(t *testing.T) {
	tests := []struct {
		name   string
		length Length
	}{
		{"Em", Em(1.5)},
		{"Rem", Rem(2)},
		{"Ch", Ch(40)},
		{"Vw", Vw(100)},
		{"Vh", Vh(50)},
		{"Cqw", Cqw(100)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.length.ToPx()
			if err == nil {
				t.Errorf("ToPx() should return error for relative unit %s", tt.length.Unit)
			}
		})
	}
}

func TestContextResolveAbsolute(t *testing.T) {
	ctx := Context{} // Empty context is fine for absolute units

	tests := []struct {
		name     string
		length   Length
		expected float64
	}{
		{"96px", Px(96), 96.0},
		{"1in", In(1), 96.0},
		{"2.54cm", Cm(2.54), 96.0},
		{"72pt", Pt(72), 96.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.length.Resolve(&ctx)
			if err != nil {
				t.Fatalf("Resolve() error = %v", err)
			}
			if !almostEqual(result.Value, tt.expected) {
				t.Errorf("Resolve() = %.6fpx, want %.6fpx", result.Value, tt.expected)
			}
		})
	}
}

func TestContextResolveFontRelative(t *testing.T) {
	ctx := Context{
		FontSize:     16.0,
		RootFontSize: 16.0,
		XHeight:      8.0,
		CapHeight:    12.0,
		ChWidth:      8.0,
		IcWidth:      16.0,
		LineHeight:   24.0,
	}

	tests := []struct {
		name     string
		length   Length
		expected float64
	}{
		{"1em", Em(1), 16.0},
		{"2em", Em(2), 32.0},
		{"1rem", Rem(1), 16.0},
		{"1.5rem", Rem(1.5), 24.0},
		{"1ex", Ex(1), 8.0},
		{"1cap", Cap(1), 12.0},
		{"40ch", Ch(40), 320.0},
		{"1ic", Ic(1), 16.0},
		{"2lh", Lh(2), 48.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.length.Resolve(&ctx)
			if err != nil {
				t.Fatalf("Resolve() error = %v", err)
			}
			if !almostEqual(result.Value, tt.expected) {
				t.Errorf("Resolve() = %.6fpx, want %.6fpx", result.Value, tt.expected)
			}
		})
	}
}

func TestContextResolveViewportRelative(t *testing.T) {
	ctx := Context{
		ViewportWidth:  1920.0,
		ViewportHeight: 1080.0,
	}

	tests := []struct {
		name     string
		length   Length
		expected float64
	}{
		{"100vw", Vw(100), 1920.0},
		{"50vw", Vw(50), 960.0},
		{"100vh", Vh(100), 1080.0},
		{"50vh", Vh(50), 540.0},
		{"10vmin", Vmin(10), 108.0}, // min(1920, 1080) = 1080, 10% = 108
		{"10vmax", Vmax(10), 192.0}, // max(1920, 1080) = 1920, 10% = 192
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.length.Resolve(&ctx)
			if err != nil {
				t.Fatalf("Resolve() error = %v", err)
			}
			if !almostEqual(result.Value, tt.expected) {
				t.Errorf("Resolve() = %.6fpx, want %.6fpx", result.Value, tt.expected)
			}
		})
	}
}

func TestContextResolveContainerRelative(t *testing.T) {
	ctx := Context{
		ContainerWidth:  800.0,
		ContainerHeight: 600.0,
	}

	tests := []struct {
		name     string
		length   Length
		expected float64
	}{
		{"100cqw", Cqw(100), 800.0},
		{"50cqw", Cqw(50), 400.0},
		{"100cqh", Cqh(100), 600.0},
		{"50cqh", Cqh(50), 300.0},
		{"10cqmin", Cqmin(10), 60.0}, // min(800, 600) = 600, 10% = 60
		{"10cqmax", Cqmax(10), 80.0}, // max(800, 600) = 800, 10% = 80
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.length.Resolve(&ctx)
			if err != nil {
				t.Fatalf("Resolve() error = %v", err)
			}
			if !almostEqual(result.Value, tt.expected) {
				t.Errorf("Resolve() = %.6fpx, want %.6fpx", result.Value, tt.expected)
			}
		})
	}
}

func TestContextResolveErrorMissingContext(t *testing.T) {
	tests := []struct {
		name   string
		length Length
		ctx    Context
	}{
		{"Em without FontSize", Em(1), Context{}},
		{"Vw without ViewportWidth", Vw(100), Context{}},
		{"Cqw without ContainerWidth", Cqw(100), Context{}},
	}

	for _, tt := range tests {
		tt := tt // capture range variable
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.length.Resolve(&tt.ctx)
			if err == nil {
				t.Errorf("Resolve() should return error when context is missing")
			}
		})
	}
}

func BenchmarkToPx(b *testing.B) {
	length := In(1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = length.ToPx()
	}
}

func BenchmarkResolveAbsolute(b *testing.B) {
	length := In(1)
	ctx := Context{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = length.Resolve(&ctx)
	}
}

func BenchmarkResolveFontRelative(b *testing.B) {
	length := Em(1.5)
	ctx := Context{FontSize: 16.0}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = length.Resolve(&ctx)
	}
}

func BenchmarkResolveViewportRelative(b *testing.B) {
	length := Vw(100)
	ctx := Context{ViewportWidth: 1920.0}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = length.Resolve(&ctx)
	}
}

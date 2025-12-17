# units

[![CI](https://github.com/SCKelemen/units/actions/workflows/ci.yml/badge.svg)](https://github.com/SCKelemen/units/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/SCKelemen/units.svg)](https://pkg.go.dev/github.com/SCKelemen/units)
[![Go Report Card](https://goreportcard.com/badge/github.com/SCKelemen/units)](https://goreportcard.com/report/github.com/SCKelemen/units)
[![License: BearWare 1.0](https://img.shields.io/badge/License-BearWare%201.0-blue.svg)](https://github.com/SCKelemen/BearWare)

A comprehensive, type-safe Go implementation of the [CSS Values and Units Module Level 4](https://www.w3.org/TR/css-values-4/) specification. Perfect for layout engines, CSS parsers, rendering systems, and any application that needs to work with CSS-style measurements.

## Features

- **Complete CSS Unit Coverage**: Supports all CSS value types including lengths, angles, time, frequency, resolution, numbers, percentages, and ratios
- **Type-Safe API**: Strongly-typed units prevent mixing incompatible measurements
- **Accurate Conversions**: Implements CSS spec-compliant conversion algorithms
- **Context-Aware Resolution**: Resolves relative units (em, vh, cqw, etc.) using rendering context
- **Zero Dependencies**: Pure Go implementation with no external dependencies
- **Well Documented**: Extensive godoc comments with references to CSS specifications and MDN docs
- **Fully Tested**: Comprehensive test coverage

## Installation

```bash
go get github.com/SCKelemen/units
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/SCKelemen/units"
)

func main() {
    // Create length values
    width := units.Px(400)
    margin := units.Em(1.5)
    height := units.Vh(100)

    // Perform arithmetic operations
    doubled := width.Mul(2)
    fmt.Println(doubled) // 800.00px

    // Convert between absolute units
    inches := units.In(1)
    pixels, _ := inches.ToPx()
    fmt.Println(pixels) // 96.00px

    // Work with angles
    angle := units.Deg(90)
    radians := angle.ToRad()
    fmt.Println(radians) // 1.57rad

    // Use time values
    duration := units.Sec(2.5)
    ms := duration.ToMs()
    fmt.Println(ms) // 2500.00ms

    // Work with ratios
    aspectRatio := units.NewRatio(16, 9)
    height := aspectRatio.ApplyToWidth(1920)
    fmt.Println(height) // 1080
}
```

## Supported Unit Types

### Length Units

#### Absolute Lengths
- **px** - Pixels (anchor unit)
- **cm** - Centimeters
- **mm** - Millimeters
- **Q** - Quarter-millimeters
- **in** - Inches
- **pt** - Points (1/72 inch)
- **pc** - Picas (12 points)

#### Font-Relative Lengths
- **em, rem** - Font size (element, root)
- **ex, rex** - X-height (element, root)
- **cap, rcap** - Cap height (element, root)
- **ch, rch** - Character width of "0" (element, root)
- **ic, ric** - Ideographic character width (element, root)
- **lh, rlh** - Line height (element, root)

#### Viewport-Relative Lengths
- **vw, vh** - Viewport width/height
- **vmin, vmax** - Viewport minimum/maximum
- **vb, vi** - Viewport block/inline size
- **svw, svh, svb, svi** - Small viewport units
- **lvw, lvh, lvb, lvi** - Large viewport units
- **dvw, dvh, dvb, dvi** - Dynamic viewport units

#### Container-Relative Lengths
- **cqw, cqh** - Container query width/height
- **cqi, cqb** - Container query inline/block size
- **cqmin, cqmax** - Container query minimum/maximum

### Other Value Types

- **Angle**: deg, grad, rad, turn
- **Time**: s, ms
- **Frequency**: Hz, kHz
- **Resolution**: dpi, dpcm, dppx
- **Number**: Dimensionless values
- **Percentage**: Relative percentages
- **Integer**: Whole numbers
- **Ratio**: Aspect ratios (e.g., 16/9)

## Context-Aware Resolution

Resolve relative units to absolute pixels using a rendering context:

```go
// Create a context with rendering information
ctx := units.Context{
    FontSize:       16.0,
    RootFontSize:   16.0,
    ViewportWidth:  1920.0,
    ViewportHeight: 1080.0,
}

// Resolve relative units
emLength := units.Em(2)
px, _ := emLength.Resolve(ctx)
fmt.Println(px) // 32.00px

vwLength := units.Vw(50)
px, _ = vwLength.Resolve(ctx)
fmt.Println(px) // 960.00px
```

## Unit Conversions

### Absolute Length Conversions

```go
// Convert between absolute units
cm := units.Cm(2.54)
px, _ := cm.ToPx()
fmt.Println(px) // 96.00px

// Use generic To() method
inches := units.In(1)
pt, _ := inches.To(units.PT)
fmt.Println(pt) // 72.00pt
```

### Angle Conversions

```go
degrees := units.Deg(180)
radians := degrees.ToRad()
turns := degrees.ToTurns()
fmt.Println(radians) // 3.14rad
fmt.Println(turns)   // 0.50turn
```

### Time Conversions

```go
seconds := units.Sec(2.5)
ms := seconds.ToMs()
fmt.Println(ms) // 2500.00ms
```

## Arithmetic Operations

```go
// Same-unit operations
a := units.Px(100)
b := units.Px(50)
sum := a.Add(b)      // 150.00px
diff := a.Sub(b)     // 50.00px

// Scalar operations
doubled := a.Mul(2)  // 200.00px
half := a.Div(2)     // 50.00px

// Comparisons
isLess := a.LessThan(b)    // false
isGreater := a.GreaterThan(b) // true
```

## Use Cases

- **Layout Engines**: Building CSS-compliant layout systems
- **CSS Parsers**: Parsing and validating CSS values
- **Design Tools**: Creating design systems with precise measurements
- **Rendering Systems**: Converting units for different display contexts
- **Animation Systems**: Working with time-based animations
- **Media Queries**: Handling viewport and container queries
- **Typography**: Managing font sizes and line heights

## Documentation

Full API documentation is available at [pkg.go.dev/github.com/SCKelemen/units](https://pkg.go.dev/github.com/SCKelemen/units).

Each type includes:
- Constructor functions for easy value creation
- String() methods for CSS-compatible output
- Arithmetic operations (Add, Sub, Mul, Div)
- Comparison operations (LessThan, GreaterThan, Equals)
- Unit conversion methods
- Utility methods specific to each type

## References

This package implements the following specifications:

- [CSS Values and Units Module Level 4](https://www.w3.org/TR/css-values-4/)
- [CSS Containment Module Level 3](https://www.w3.org/TR/css-contain-3/) (Container queries)
- [MDN Web Docs - CSS values and units](https://developer.mozilla.org/en-US/docs/Web/CSS/CSS_Values_and_Units)
- [web.dev - Learn CSS](https://web.dev/learn/css/)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

## License

BearWare 1.0 License - see the [LICENSE](LICENSE) file for details.

## Origin

Originally implemented in [github.com/SCKelemen/layout](https://github.com/SCKelemen/layout) and extracted as a standalone package for reuse across layout engines, text rendering, and other CSS-based projects.

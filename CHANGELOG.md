# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Planned
- CSS calc() expression support
- CSS color value types
- Additional utility functions based on user feedback

## [1.0.2] - 2025-12-17

### Fixed
- Fixed all remaining golangci-lint errors
- Fixed `Clamp()` parameter shadowing in Integer, Number, and Percentage types (renamed `min`/`max` to `minValue`/`maxValue`)
- Fixed G601 implicit memory aliasing in test loop
- Renamed `Integer.Min()`/`Max()` methods to `Minimum()`/`Maximum()` to avoid confusion with Go 1.21+ built-in functions

### Changed
- Updated CI to test on Go 1.21, 1.23, and 1.25 (replaced 1.22 with 1.25)
- All lint, vet, and build checks now use Go 1.25

### Breaking Changes
- `Integer.Min()` is now `Integer.Minimum()`
- `Integer.Max()` is now `Integer.Maximum()`

## [1.0.1] - 2025-12-17

### Fixed
- Fixed golangci-lint errors for built-in identifier shadowing (min/max variables)
- Changed Context parameter to pointer to reduce memory overhead
- Removed deprecated golangci-lint configuration options
- Updated documentation examples to reflect pointer usage

### Changed
- `Length.Resolve(ctx Context)` is now `Length.Resolve(ctx *Context)` (breaking change for type signature)

## [1.0.0] - 2025-12-17

### Added
- Initial release of CSS Values and Units Module Level 4 implementation
- Complete type-safe API for CSS value types:
  - **Length**: All absolute, font-relative, viewport-relative, and container-relative units
  - **Angle**: deg, grad, rad, turn
  - **Time**: s, ms
  - **Frequency**: Hz, kHz
  - **Resolution**: dpi, dpcm, dppx
  - **Number**: Dimensionless numeric values
  - **Percentage**: Relative percentage values
  - **Integer**: Whole number values
  - **Ratio**: Aspect ratios (e.g., 16:9)
- Unit conversion methods for all types
- Context-aware resolution for relative units
- Arithmetic operations (Add, Sub, Mul, Div)
- Comparison operations (LessThan, GreaterThan, Equals)
- Comprehensive test coverage (73.8%)
- Full godoc documentation with CSS spec references
- Zero external dependencies
- GitHub Actions CI workflow
- Comprehensive README with examples

### Documentation
- Package-level documentation
- CSS Values Level 4 specification references
- MDN Web Docs references
- web.dev learning resource references
- Example code for all major features

## Origin

This package was originally implemented in [github.com/SCKelemen/layout](https://github.com/SCKelemen/layout) and extracted as a standalone package for reuse across layout engines, text rendering, and other CSS-based projects.

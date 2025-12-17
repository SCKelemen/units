# Release Checklist for v1.0.0

This checklist ensures the package is ready for v1.0.0 release.

## Pre-Release Checklist

### Code Quality
- [x] All tests passing (`go test ./...`)
- [x] Race detector clean (`go test -race ./...`)
- [x] Code coverage >70% (currently 73.8%)
- [x] `go vet` passes with no errors
- [x] Code is properly formatted (`gofmt -w .`)
- [ ] Run `golangci-lint run` locally (optional, will run in CI)

### Documentation
- [x] Package-level documentation is complete (length.go)
- [x] All public types have godoc comments
- [x] All public functions have godoc comments with examples
- [x] References to CSS specs added to all relevant files
- [x] References to MDN documentation added
- [x] README.md is comprehensive with:
  - [x] Installation instructions
  - [x] Quick start examples
  - [x] Feature list
  - [x] API overview
  - [x] Use cases
  - [x] Badges (CI, GoDoc, Go Report Card, License)
- [x] LICENSE file added (BearWare 1.0)

### Repository Setup
- [x] GitHub Actions CI workflow configured (.github/workflows/ci.yml)
  - [x] Testing on Go 1.21, 1.23, 1.25
- [x] golangci-lint configuration added (.golangci.yml)
- [x] go.mod specifies Go 1.21+ compatibility
- [x] .gitignore includes common Go artifacts

### Testing
- [x] Unit tests for all types
- [x] Conversion tests
- [x] Arithmetic operation tests
- [x] Comparison operation tests
- [x] Edge case tests (zero values, negative values, etc.)
- [x] Context resolution tests

## Release Steps

### 1. Final Code Review
- [ ] Review all code for potential breaking changes
- [ ] Verify API is clean and intuitive
- [ ] Ensure backward compatibility is maintained for future releases
- [ ] Check for any TODO comments or incomplete features

### 2. Version Tagging
- [ ] Update CHANGELOG.md (create if needed)
- [ ] Commit all final changes
- [ ] Tag the release: `git tag v1.0.0`
- [ ] Push the tag: `git push origin v1.0.0`

### 3. GitHub Release
- [ ] Create GitHub release from tag
- [ ] Include release notes highlighting:
  - Complete CSS Values Level 4 implementation
  - All supported unit types
  - Key features (type-safety, conversions, context-aware resolution)
  - Zero dependencies
- [ ] Attach any relevant documentation

### 4. Go Module Registry
- [ ] Verify module is discoverable at pkg.go.dev
- [ ] Check that documentation renders correctly on pkg.go.dev
- [ ] Verify examples show up in documentation

### 5. Post-Release
- [ ] Announce release (if applicable)
- [ ] Update any dependent projects
- [ ] Monitor for issues or bug reports
- [ ] Set up issue templates for bug reports and feature requests

## v1.0.0 API Stability Guarantee

Once v1.0.0 is released:
- No breaking changes to public API
- Follow semantic versioning (semver)
- Bug fixes: patch version (v1.0.1)
- New features: minor version (v1.1.0)
- Breaking changes: major version (v2.0.0)

## Notes

### What's Included in v1.0.0
- Complete implementation of CSS Values Level 4 specification
- All CSS unit types: Length, Angle, Time, Frequency, Resolution, Number, Percentage, Integer, Ratio
- Absolute and relative length units
- Font-relative, viewport-relative, and container-relative units
- Type-safe API with unit checking
- Context-aware resolution for relative units
- Comprehensive test coverage
- Zero external dependencies

### Future Enhancements (Post v1.0.0)
- CSS calc() expression support
- CSS color types
- Additional utility methods based on user feedback
- Performance optimizations
- Extended examples and tutorials

## Testing the Release

Before tagging v1.0.0, test the package in a separate project:

```bash
# Create a test project
mkdir test-units && cd test-units
go mod init test-units

# Add the package
go get github.com/SCKelemen/units@main

# Write a simple test program
cat > main.go << 'EOF'
package main

import (
    "fmt"
    "github.com/SCKelemen/units"
)

func main() {
    // Test various units
    px := units.Px(96)
    inch, _ := px.To(units.IN)
    fmt.Printf("%v = %v\n", px, inch)

    angle := units.Deg(180)
    fmt.Printf("%v = %v\n", angle, angle.ToRad())

    time := units.Sec(2.5)
    fmt.Printf("%v = %v\n", time, time.ToMs())

    ratio := units.NewRatio(16, 9)
    fmt.Printf("16:9 ratio applied to 1920px width = %vpx height\n", ratio.ApplyToWidth(1920))
}
EOF

# Run it
go run main.go
```

## Sign-Off

- [ ] Project maintainer review
- [ ] All checklist items completed
- [ ] Ready for v1.0.0 release

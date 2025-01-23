//go:build ios && cgo

package main

// Force CGO linking for iOS builds
import "C"

// This file ensures proper build tags are set for iOS builds
const (
	// Build tags that should be included
	_ = "cgo"
	_ = "ios"
)

func init() {
	// This empty init function ensures the file is included in the build
}

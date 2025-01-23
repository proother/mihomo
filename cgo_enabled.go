//go:build ios

package main

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation

#import <Foundation/Foundation.h>

void ensureCGOEnabled() {
    @autoreleasepool {
        [[NSProcessInfo processInfo] operatingSystemVersion];
    }
}
*/
import "C"

func init() {
	// Call C function to ensure CGO is properly enabled and linked
	C.ensureCGOEnabled()
}

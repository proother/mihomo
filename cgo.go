//go:build ios && cgo

package main

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation -framework CoreFoundation -framework Security -framework SystemConfiguration
#import <Foundation/Foundation.h>

void initializeRuntime() {
    @autoreleasepool {
        [NSThread new];
    }
}
*/
import "C"

func init() {
	C.initializeRuntime()
}

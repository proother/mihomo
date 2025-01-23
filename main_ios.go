//go:build ios

package main

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation -framework CoreFoundation -framework Security -framework SystemConfiguration
*/
import "C"

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/metacubex/mihomo/constant"
	"github.com/metacubex/mihomo/constant/features"
	"github.com/metacubex/mihomo/hub"
	"github.com/metacubex/mihomo/log"
)

func init() {
	// iOS-specific initialization
	if err := hub.InitializeHub(); err != nil {
		log.Fatalln("Failed to initialize hub:", err)
	}
}

func main() {
	fmt.Printf("Mihomo Meta %s %s %s with %s %s\n",
		constant.Version, runtime.GOOS, runtime.GOARCH, runtime.Version(), constant.BuildTime)
	if tags := features.Tags(); len(tags) != 0 {
		fmt.Printf("Use tags: %s\n", strings.Join(tags, ", "))
	}

	// iOS-specific main loop
	select {} // Keep the process running
}

//go:build ios

package libbox

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation
#import <Foundation/Foundation.h>

const char* getHomeDirectory() {
	NSString *path = NSHomeDirectory();
	return [path UTF8String];
}
*/
import "C"
import (
	"fmt"
	"path/filepath"
	"unsafe"
)

// Error mapping
var errorMap = map[C.TunnelError]string{
	C.TunnelErrorNone:                 "no error",
	C.TunnelErrorPermissionDenied:     "permission denied",
	C.TunnelErrorNetworkError:         "network error",
	C.TunnelErrorSystemError:          "system error",
	C.TunnelErrorInvalidConfiguration: "invalid configuration",
	C.TunnelErrorTunnelStartFailed:    "tunnel start failed",
	C.TunnelErrorDNSConfigFailed:      "DNS configuration failed",
	C.TunnelErrorProxyConfigFailed:    "proxy configuration failed",
	C.TunnelErrorMemoryError:          "memory error",
	C.TunnelErrorTimeout:              "operation timed out",
	C.TunnelErrorAlreadyRunning:       "tunnel already running",
}

// Enhanced error type
type PlatformError struct {
	Code    C.TunnelError
	Message string
}

func (e *PlatformError) Error() string {
	if msg, ok := errorMap[e.Code]; ok {
		return fmt.Sprintf("%s: %s", msg, e.Message)
	}
	return fmt.Sprintf("unknown error (%d): %s", e.Code, e.Message)
}

// Initialize platform with error handling
func initializePlatform() error {
	if code := C.initializeNetworkExtension(); code != C.TunnelErrorNone {
		return &PlatformError{
			Code:    code,
			Message: "failed to initialize network extension",
		}
	}
	return nil
}

// Setup tunnel with enhanced error handling
func setupTunnelDevice(name string, mtu int) error {
	if name == "" {
		return &PlatformError{
			Code:    C.TunnelErrorInvalidConfiguration,
			Message: "tunnel name cannot be empty",
		}
	}

	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	if code := C.setupTunnel(cName); code != C.TunnelErrorNone {
		return &PlatformError{
			Code:    code,
			Message: fmt.Sprintf("failed to setup tunnel: %s", name),
		}
	}
	return nil
}

// Configure DNS settings
func configureDNS(servers []string) error {
	if len(servers) == 0 {
		return nil
	}

	// Convert DNS servers to comma-separated string
	dnsServers := ""
	for i, server := range servers {
		if i > 0 {
			dnsServers += ","
		}
		dnsServers += server
	}

	cServers := C.CString(dnsServers)
	defer C.free(unsafe.Pointer(cServers))

	if code := C.configureDNS(cServers); code != C.TunnelErrorNone {
		return &PlatformError{
			Code:    code,
			Message: "failed to configure DNS",
		}
	}
	return nil
}

// Set system proxy settings
func setSystemProxy(host string, port int) error {
	cHost := C.CString(host)
	defer C.free(unsafe.Pointer(cHost))

	if code := C.setSystemProxy(cHost, C.int(port)); code != C.TunnelErrorNone {
		return &PlatformError{
			Code:    code,
			Message: "failed to set system proxy",
		}
	}
	return nil
}

func GetHomeDir() string {
	cstr := C.getHomeDirectory()
	return C.GoString(cstr)
}

func GetConfigDir() string {
	return filepath.Join(GetHomeDir(), "Documents")
}

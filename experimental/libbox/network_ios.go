//go:build ios

package libbox

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation -framework CoreFoundation -framework Security -framework SystemConfiguration
#import <Foundation/Foundation.h>
#import <SystemConfiguration/SystemConfiguration.h>
*/
import "C"

import (
	"net"
	"net/netip"
)

type NetworkInfo struct {
	Interface string
	Address   netip.Addr
}

func GetNetworkInfo() (*NetworkInfo, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	for _, iface := range ifaces {
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok {
				if ip, err := netip.ParseAddr(ipnet.IP.String()); err == nil {
					return &NetworkInfo{
						Interface: iface.Name,
						Address:   ip,
					}, nil
				}
			}
		}
	}

	return nil, nil
}

func setupNetworkExtension() error {
	C.setupPacketTunnelProvider()
	return nil
}

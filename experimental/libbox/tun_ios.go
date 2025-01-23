//go:build ios

package libbox

import (
	"errors"
)

type TUN struct {
	device string
	mtu    int
}

func NewSystemTUN(name string, mtu int) (*TUN, error) {
	if name == "" {
		return nil, errors.New("empty tun device name")
	}
	if mtu <= 0 {
		mtu = 1500
	}
	return &TUN{
		device: name,
		mtu:    mtu,
	}, nil
}

func (t *TUN) Name() string {
	return t.device
}

func (t *TUN) MTU() int {
	return t.mtu
}

func (t *TUN) Read(buff []byte) (int, error) {
	// Implement iOS-specific TUN read
	return 0, nil
}

func (t *TUN) Write(buff []byte) (int, error) {
	// Implement iOS-specific TUN write
	return 0, nil
}

func (t *TUN) Close() error {
	// Implement iOS-specific cleanup
	return nil
}

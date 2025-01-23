//go:build ios

package libbox

import (
	tun "github.com/metacubex/sing-tun"
)

func createTUN(options *TUNOptions) (tun.Tuple, error) {
	return tun.NewSystemTUN(options.Name, options.MTU)
}

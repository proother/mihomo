package libbox

import (
	"encoding/json"

	"github.com/metacubex/mihomo/experimental/libbox/option"
)

// LogOptions represents logging configuration
type LogOptions struct {
	Level string `json:"level"`
}

// DNSOptions represents DNS configuration
type DNSOptions = option.DNS

// TUNOptions represents TUN device configuration
type TUNOptions = option.TUN

// InboundOptions represents inbound configuration
type InboundOptions = option.Inbound

// Config represents the configuration structure
type Config struct {
	LogLevel string           `json:"log_level"`
	DNS      *option.DNS      `json:"dns"`
	TUN      *option.TUN      `json:"tun"`
	Inbound  []option.Inbound `json:"inbound"`
}

func (c *Box) SetConfig(config string) error {
	var opts Config
	if err := json.Unmarshal([]byte(config), &opts); err != nil {
		return err
	}

	if opts.DNS != nil {
		c.SetDNS(*opts.DNS)
	}
	if opts.TUN != nil {
		c.SetTUN(*opts.TUN)
	}
	if opts.Inbound != nil {
		c.SetInbound(opts.Inbound)
	}

	return validateOptions(c.options)
}

func (c *Config) ToOptions() *option.Options {
	return &option.Options{
		DNS:     c.DNS,
		TUN:     c.TUN,
		Inbound: c.Inbound,
	}
}

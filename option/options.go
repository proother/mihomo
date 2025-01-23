package option

// RouteOptions defines routing configuration
type RouteOptions struct {
	// Add necessary fields
}

// ExperimentalOptions defines experimental features
type ExperimentalOptions struct {
	// Add necessary fields
}

// FakeIPOptions defines fake IP configuration
type FakeIPOptions struct {
	// Add necessary fields
}

// Options contains all options
type Options struct {
	LogLevel string
	DNS      *DNSOptions
	TUN      *TUNOptions
	Inbound  *InboundOptions
}

// LogConfig defines logging configuration
type LogConfig struct {
	Level string
}

// DNSOptions defines DNS configuration
type DNSOptions struct {
	Enable   bool           `json:"enable,omitempty"`
	Servers  []string       `json:"servers,omitempty"`
	Strategy string         `json:"strategy,omitempty"`
	FakeIP   *FakeIPOptions `json:"fakeip,omitempty"`
}

// TUNOptions defines TUN interface configuration
type TUNOptions struct {
	Enable bool
	Device string
	MTU    int
	Check  func() error
}

// InboundOptions defines inbound configuration
type InboundOptions struct {
	Enable bool
	// Add other inbound fields
}

// OutboundOptions defines outbound proxy options
type OutboundOptions struct {
	Type   string `json:"type"`
	Tag    string `json:"tag,omitempty"`
	Server string `json:"server,omitempty"`
	Port   int    `json:"port,omitempty"`
	// Type-specific options follow
}

// Example usage in config.json:
/*
{
    "log": {
        "level": "info",
        "timestamp": true
    },
    "dns": {
        "enable": true,
        "servers": ["8.8.8.8", "1.1.1.1"]
    },
    "tun": {
        "enable": true,
        "device": "utun0",
        "mtu": 9000
    },
    "inbounds": [
        {
            "type": "mixed",
            "listen": "127.0.0.1",
            "port": 7890
        }
    ],
    "outbounds": [
        {
            "type": "direct",
            "tag": "direct"
        },
        {
            "type": "block",
            "tag": "block"
        }
    ]
}
*/

// Add validation methods
func (o *Options) Validate() error {
	if o.TUN != nil {
		// Add validation logic
	}
	return nil
}

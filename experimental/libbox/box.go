package libbox

import (
	"context"
	"sync"

	"github.com/metacubex/mihomo/experimental/libbox/option"
)

// Box represents the main application container
type Box struct {
	ctx      context.Context
	cancel   context.CancelFunc
	mutex    sync.Mutex
	running  bool
	options  *option.Options
	instance any
}

// Instance represents a running instance
type Instance interface {
	// Add necessary instance methods
}

// NewBox creates a new Box instance
func NewBox(options *option.Options) (*Box, error) {
	if err := validateOptions(options); err != nil {
		return nil, err
	}

	box := &Box{
		options: options,
	}
	return box, nil
}

func validateOptions(options *option.Options) error {
	if options == nil {
		return nil
	}
	if options.TUN != nil && options.TUN.Enable && options.TUN.Check != nil {
		return options.TUN.Check()
	}
	if options.DNS != nil && options.DNS.Enable && options.DNS.Check != nil {
		return options.DNS.Check()
	}
	return nil
}

func (b *Box) Start() error {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if b.running {
		return nil
	}

	// Initialize instance
	instance, err := NewInstance(b.options)
	if err != nil {
		return err
	}

	b.instance = instance
	b.running = true
	return nil
}

func (b *Box) Stop() error {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if !b.running {
		return nil
	}

	// Cleanup and stop
	if b.cancel != nil {
		b.cancel()
	}

	b.running = false
	return nil
}

func (b *Box) IsRunning() bool {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	return b.running
}

func (b *Box) initializeTUN(tun *option.TUN) error {
	if tun == nil || !tun.Enable {
		return nil
	}
	// Implementation
	return nil
}

func (b *Box) initializeDNS(dns *option.DNS) error {
	if dns == nil || !dns.Enable {
		return nil
	}
	// Implementation
	return nil
}

// NewInstance creates a new Box instance
func NewInstance(options *option.Options) (*Box, error) {
	if err := validateOptions(options); err != nil {
		return nil, err
	}

	box := &Box{
		options: options,
	}
	return box, nil
}

// SetDNS updates the DNS configuration
func (b *Box) SetDNS(dns option.DNS) {
	if b.options == nil {
		b.options = &option.Options{}
	}
	b.options.DNS = &dns
}

// SetTUN updates the TUN configuration
func (b *Box) SetTUN(tun option.TUN) {
	if b.options == nil {
		b.options = &option.Options{}
	}
	b.options.TUN = &tun
}

// SetInbound updates the inbound configuration
func (b *Box) SetInbound(inbound []option.Inbound) {
	if b.options == nil {
		b.options = &option.Options{}
	}
	b.options.Inbound = inbound
}

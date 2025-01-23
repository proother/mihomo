package option

// CheckFunc is a function type for configuration validation
type CheckFunc func() error

// DNS configuration options
type DNS struct {
	Enable            bool      `yaml:"enable"`
	Listen            string    `yaml:"listen"`
	NameServers       []string  `yaml:"nameservers"`
	EnhancedMode      string    `yaml:"enhanced-mode"`
	FallbackFilter    bool      `yaml:"fallback-filter"`
	DefaultNameserver []string  `yaml:"default-nameserver"`
	Check             CheckFunc `yaml:"-"`
}

// TUN configuration options
type TUN struct {
	Enable     bool      `yaml:"enable"`
	Device     string    `yaml:"device"`
	Stack      string    `yaml:"stack"`
	AutoRoute  bool      `yaml:"auto-route"`
	AutoDetect bool      `yaml:"auto-detect"`
	DNS        []string  `yaml:"dns"`
	Check      CheckFunc `yaml:"-"`
}

// Inbound configuration options
type Inbound struct {
	Type        string            `yaml:"type"`
	Listen      string            `yaml:"listen"`
	Port        int               `yaml:"port"`
	Settings    map[string]string `yaml:"settings"`
	TLS         bool              `yaml:"tls"`
	Certificate string            `yaml:"certificate"`
}

// Options represents the main configuration options
type Options struct {
	DNS     *DNS      `yaml:"dns"`
	TUN     *TUN      `yaml:"tun"`
	Inbound []Inbound `yaml:"inbound"`
}

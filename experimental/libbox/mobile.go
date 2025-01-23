package libbox

// MobileConfig represents mobile configuration
type MobileConfig struct {
	ConfigPath string
	WorkingDir string
}

func (c *MobileConfig) GetConfigPath() string {
	return c.ConfigPath
}

func (c *MobileConfig) GetWorkingDir() string {
	return c.WorkingDir
}

// NewConfig creates a new mobile configuration
func NewConfig(configPath, workingDir string) *MobileConfig {
	return &MobileConfig{
		ConfigPath: configPath,
		WorkingDir: workingDir,
	}
}

//export StartMihomo
func StartMihomo(config *MobileConfig) error {
	return startService(config)
}

//export StopMihomo
func StopMihomo() error {
	return stopService()
}

// StartService starts the service with given options
func StartService(opts *MobileConfig) error {
	// Implementation
	return nil
}

func startService(opts *MobileConfig) error {
	// Implementation
	return nil
}

func stopService() error {
	// Implementation
	return nil
}

package modules

import "github.com/gjh33/GoNativeECS/examples/battlesim/systems"

// InitializeModule resets the battle simulation and prepares it to run
type InitializeModule struct{}

// NewInitializeModule returns a default constructed initialization module
func NewInitializeModule() *InitializeModule {
	return new(InitializeModule)
}

// Perform implements Module interface
func (module *InitializeModule) Perform() {
	systems.Battle().ResetSimulation()
}

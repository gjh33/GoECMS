package modules

import (
	"github.com/gjh33/GoNativeECS/examples/battlesim/systems"
)

// SimulateRoundModule will simulate one round of fighting from the battle system
type SimulateRoundModule struct {
	FightsPerRound int
}

// NewSimulateRoundModule is the default constructor for this module
func NewSimulateRoundModule() (obj *SimulateRoundModule) {
	obj = new(SimulateRoundModule)
	obj.FightsPerRound = 1
	return
}

// Perform implements Module interface
func (module *SimulateRoundModule) Perform() {
	for i := 0; i < module.FightsPerRound; i++ {
		systems.Battle().SimulateRandomFight()
	}
}

package examples

import (
	"fmt"

	"github.com/gjh33/GoNativeECS/ecms"
	"github.com/gjh33/GoNativeECS/examples/battlesim/components"
	modules "github.com/gjh33/GoNativeECS/examples/battlesim/modules/battle"
	"github.com/gjh33/GoNativeECS/examples/battlesim/systems"
)

// BattleSimWarriorCount tells the battle sim example how many warriors to simulate in battle
var BattleSimWarriorCount = 100

// BasicWarrior is an Entity
type BasicWarrior struct {
	*ecms.BasicEntity
	*components.Health
	*components.Weapon
	*components.BattleStats
}

// CreateHuman creates human type entity
func CreateHuman() (obj *BasicWarrior) {
	obj = new(BasicWarrior)
	obj.BasicEntity = ecms.NewBasicEntity()
	obj.Health = &components.Health{
		Current: 100,
		Max:     100,
	}
	obj.Weapon = &components.Weapon{
		Name:   "Sword",
		Range:  10,
		Damage: 7,
	}
	obj.BattleStats = &components.BattleStats{
		FightCount:          0,
		TotalDamageDealt:    0,
		TotalDamageReceived: 0,
	}
	return
}

// RunBattleSim runs the simulation to completion
func RunBattleSim() {
	env := new(ecms.Environment)
	env.AddSystem(systems.Battle())

	for i := 0; i < BattleSimWarriorCount; i++ {
		env.AddEntity(CreateHuman())
	}

	initMod := modules.NewInitializeModule()
	initMod.Perform()

	updateMod := modules.NewSimulateRoundModule()

	for len(systems.Battle().Alive) > 1 {
		updateMod.Perform()
	}

	placements := 5
	fmt.Printf("Top %v:\n", placements)
	for ind, warrior := range systems.Battle().Dead[len(systems.Battle().Dead)-placements : len(systems.Battle().Dead)] {
		place := placements - ind
		strct, ok := warrior.(*BasicWarrior)
		if ok {
			fmt.Printf("#%v: %+v\n", place, strct.UID())
			fmt.Printf("\t%+v\n", strct.BattleStatsComponent())
		}
	}
}

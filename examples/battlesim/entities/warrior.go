package entities

import (
	"github.com/gjh33/GoNativeECS/ecms"
	"github.com/gjh33/GoNativeECS/examples/battlesim/components"
)

// WarriorEntity is an entity capable of fighting
type WarriorEntity interface {
	ecms.Entity
	components.HealthComponent
	components.WeaponComponent
	components.BattleStatsComponent
}

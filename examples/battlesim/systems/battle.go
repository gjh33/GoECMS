package systems

import (
	"math/rand"
	"time"

	"github.com/gjh33/GoNativeECS/ecms"
	"github.com/gjh33/GoNativeECS/examples/battlesim/entities"
)

var instance *BattleSystem

// BattleSystem simulates random battles until one is left
type BattleSystem struct {
	// Entities Accepted
	Warriors []entities.WarriorEntity
	// System State
	Alive  []entities.WarriorEntity
	Dead   []entities.WarriorEntity
	Fights int
	// Private
	rng *rand.Rand
}

// Accept implements the System interfrace.
// it takes warrior entities to operate on.
func (system *BattleSystem) Accept(entity ecms.Entity) {
	warrior, ok := entity.(entities.WarriorEntity)
	if ok {
		system.Warriors = append(system.Warriors, warrior)
	}
}

// NewBattleSystem is the constructor for a battle system
func NewBattleSystem() (obj *BattleSystem) {
	obj = new(BattleSystem)
	obj.Warriors = make([]entities.WarriorEntity, 0)
	obj.Alive = make([]entities.WarriorEntity, 0)
	obj.Dead = make([]entities.WarriorEntity, 0)
	obj.rng = rand.New(rand.NewSource(time.Now().Unix()))
	return
}

// Battle returns the singleton instance of a system and if none exists makes one
func Battle() *BattleSystem {
	if instance == nil {
		instance = NewBattleSystem()
	}
	return instance
}

// SYSTEM METHODS

// ResetSimulation recovers all warrior's health and sets them all to alive
func (system *BattleSystem) ResetSimulation() {
	// Reset the slices and underlying arrays
	// Since we know an exact cap on the size, and any simulation running to completion will
	// fill up both slices at one point, we should just automatically size to number of warriors
	system.Alive = make([]entities.WarriorEntity, 0, len(system.Warriors))
	system.Alive = make([]entities.WarriorEntity, 0, len(system.Warriors))

	for _, warrior := range system.Warriors {
		if warrior.Enabled() {
			warrior.Health().Current = warrior.Health().Max
			system.Alive = append(system.Alive, warrior)
		}
	}
}

// SimulateRandomFight picks two random alive warriors and makes them fight
func (system *BattleSystem) SimulateRandomFight() {
	w1, w2 := system.pickTwoAlive()

	// Update Data
	w1.Health().Current -= w2.Weapon().Damage
	w2.Health().Current -= w1.Weapon().Damage
	w1.BattleStats().FightCount++
	w2.BattleStats().FightCount++
	w1.BattleStats().TotalDamageDealt += w1.Weapon().Damage
	w2.BattleStats().TotalDamageDealt += w2.Weapon().Damage
	w1.BattleStats().TotalDamageReceived += w2.Weapon().Damage
	w2.BattleStats().TotalDamageReceived += w1.Weapon().Damage

	// System internals
	if w1.Health().Current <= 0 {
		system.kill(w1)
	}
	if w2.Health().Current <= 0 {
		system.kill(w2)
	}
}

// PRIVATE METHODS
func (system *BattleSystem) pickTwoAlive() (entities.WarriorEntity, entities.WarriorEntity) {
	ind1 := system.rng.Intn(len(system.Alive))
	// For second index, roll with one less max value.
	// If number rolled is >= ind1, then increment it.
	// This ensures a unique pair
	ind2 := system.rng.Intn(len(system.Alive) - 1)
	if ind2 >= ind1 {
		ind2++
	}

	return system.Alive[ind1], system.Alive[ind2]
}

func (system *BattleSystem) kill(warrior entities.WarriorEntity) {
	ind := -1
	for i, w := range system.Alive {
		if w == warrior {
			ind = i
			break
		}
	}

	// If we failed to find it, it wasn't alive anyways
	if ind < 0 {
		return
	}

	// Delete it from the list
	copy(system.Alive[ind:], system.Alive[ind+1:])
	system.Alive[len(system.Alive)-1] = nil
	system.Alive = system.Alive[:len(system.Alive)-1]

	system.Dead = append(system.Dead, warrior)
}

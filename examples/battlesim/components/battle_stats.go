package components

// BattleStatsComponent is used to identify that an entity has battle stats
type BattleStatsComponent interface {
	BattleStatsComponent() *BattleStats
}

// BattleStats represents this entity's history in battle
type BattleStats struct {
	FightCount          int
	TotalDamageDealt    float32
	TotalDamageReceived float32
}

// BattleStatsComponent implements BattleStatsComponent
func (stats *BattleStats) BattleStatsComponent() *BattleStats {
	return stats
}

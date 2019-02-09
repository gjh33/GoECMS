package components

// HealthComponent is used to identify that an entity has health
type HealthComponent interface {
	Health() *Health
}

// Health is the data behind the health component
type Health struct {
	Max     float32
	Current float32
}

// Health implements HealthComponent interface
func (health *Health) Health() *Health {
	return health
}

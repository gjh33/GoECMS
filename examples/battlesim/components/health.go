package components

// HealthComponent is used to identify that an entity has health
type HealthComponent interface {
	HealthComponent() *Health
}

// Health is the data behind the health component
type Health struct {
	Max     float32
	Current float32
}

// HealthComponent implements HealthComponent interface
func (health *Health) HealthComponent() *Health {
	return health
}

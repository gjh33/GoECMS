package ecms

// Environment represents a set of systems and entities. It manages the injection of entities into systems etc.
// entities are added to systems as they arrive, and when a system is added all existing compatible entities are added
// to that system
type Environment struct {
	systems  []System
	entities []Entity
}

// AddSystem adds the system to the environment and checks for compatible entities
func (env *Environment) AddSystem(sys System) {
	env.systems = append(env.systems, sys)
	for _, entity := range env.entities {
		sys.Accept(entity)
	}
}

// AddEntity adds the entity to the environment and adds it to any compatible systems
func (env *Environment) AddEntity(entity Entity) {
	env.entities = append(env.entities, entity)
	for _, sys := range env.systems {
		sys.Accept(entity)
	}
}

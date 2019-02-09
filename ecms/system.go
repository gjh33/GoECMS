package ecms

// System defines the required interface for a system that acts on
type System interface {
	// Accept will be passed components and should use go native type asserts to add these components to a group
	Accept(Entity)
}
